// Copyright 2024 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package service

import (
	"context"
	"time"

	"github.com/juju/clock"
	"github.com/juju/collections/set"
	"github.com/juju/errors"
	"github.com/juju/worker/v4"
	"github.com/juju/worker/v4/catacomb"
	"gopkg.in/tomb.v2"

	"github.com/juju/juju/core/changestream"
	"github.com/juju/juju/core/leadership"
	"github.com/juju/juju/core/secrets"
	"github.com/juju/juju/core/watcher"
	"github.com/juju/juju/core/watcher/watchertest"
	"github.com/juju/juju/internal/secrets/provider"
)

// WatchableService provides the API for working with the secret service.
type WatchableService struct {
	SecretService
	watcherFactory WatcherFactory
}

// NewWatchableService returns a new watchable service wrapping the specified state.
func NewWatchableService(
	st State, logger Logger, watcherFactory WatcherFactory, adminConfigGetter BackendAdminConfigGetter,
) *WatchableService {
	return &WatchableService{
		SecretService: SecretService{
			st:                st,
			logger:            logger,
			clock:             clock.WallClock,
			providerGetter:    provider.Provider,
			adminConfigGetter: adminConfigGetter,
		},
		watcherFactory: watcherFactory,
	}
}

// WatchConsumedSecretsChanges watches secrets consumed by the specified unit
// and returns a watcher which notifies of secret URIs that have had a new revision added.
func (s *WatchableService) WatchConsumedSecretsChanges(ctx context.Context, unitName string) (watcher.StringsWatcher, error) {
	ch := make(chan []string, 1)
	ch <- []string{}
	return watchertest.NewMockStringsWatcher(ch), nil
}

// WatchRemoteConsumedSecretsChanges watches secrets remotely consumed by any unit
// of the specified app and retuens a watcher which notifies of secret URIs
// that have had a new revision added.
func (s *WatchableService) WatchRemoteConsumedSecretsChanges(ctx context.Context, appName string) (watcher.StringsWatcher, error) {
	ch := make(chan []string, 1)
	ch <- []string{}
	return watchertest.NewMockStringsWatcher(ch), nil
}

// WatchObsolete returns a watcher for notifying when:
//   - a secret owned by the entity is deleted
//   - a secret revision owned by the entity no longer
//     has any consumers
//
// Obsolete revisions results are "uri/revno" and deleted
// secret results are "uri".
func (s *WatchableService) WatchObsolete(ctx context.Context, owners ...CharmSecretOwner) (watcher.StringsWatcher, error) {
	if len(owners) == 0 {
		return nil, errors.New("at least one owner must be provided")
	}

	appOwners, unitOwners := splitCharmSecretOwners(owners...)
	table, query := s.st.InitialWatchStatementForObsoleteRevision(ctx, appOwners, unitOwners)
	w, err := s.watcherFactory.NewNamespaceWatcher(
		table, changestream.Update, query,
	)
	if err != nil {
		return nil, errors.Trace(err)
	}
	processChanges := func(ctx context.Context, revisionUUIDs ...string) ([]string, error) {
		return s.st.GetRevisionIDsForObsolete(ctx, appOwners, unitOwners, revisionUUIDs...)
	}
	return newObsoleteRevisionWatcher(w, s.logger, processChanges)
}

// obsoleteRevisionWatcher watches for obsolete changes to secret revisions.
// It receives a stream of revision UUIDs and processes them to determine which
// are obsolete. It then sends the slice of `<uri>/<version>` to its output channel.
type obsoleteRevisionWatcher struct {
	catacomb catacomb.Catacomb
	logger   Logger

	sourceWatcher  watcher.StringsWatcher
	processChanges func(
		ctx context.Context,
		revisionUUID ...string,
	) ([]string, error)

	out chan []string
}

func newObsoleteRevisionWatcher(
	sourceWatcher watcher.StringsWatcher, logger Logger,
	processChanges func(ctx context.Context, revisionUUID ...string) ([]string, error),
) (*obsoleteRevisionWatcher, error) {
	w := &obsoleteRevisionWatcher{
		sourceWatcher:  sourceWatcher,
		logger:         logger,
		processChanges: processChanges,
		out:            make(chan []string),
	}
	err := catacomb.Invoke(catacomb.Plan{
		Site: &w.catacomb,
		Work: w.loop,
		Init: []worker.Worker{sourceWatcher},
	})
	return w, errors.Trace(err)
}

func (w *obsoleteRevisionWatcher) getRevisionIDs(revisionUUIDs ...string) ([]string, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return w.processChanges(w.catacomb.Context(ctx), revisionUUIDs...)
}

func (w *obsoleteRevisionWatcher) loop() error {
	defer close(w.out)

	var changes set.Strings
	// To allow the initial event to be sent.
	out := w.out

	addChanges := func(processed ...string) {
		if len(processed) == 0 {
			return
		}
		if changes == nil {
			changes = set.NewStrings()
		}
		for _, change := range processed {
			changes.Add(change)
		}
		out = w.out
	}

	for {
		select {
		case <-w.catacomb.Dying():
			return w.catacomb.ErrDying()
		case revisionUUIDs, ok := <-w.sourceWatcher.Changes():
			if !ok {
				return errors.Errorf("event watcher closed")
			}
			if len(revisionUUIDs) == 0 {
				continue
			}
			processed, err := w.getRevisionIDs(revisionUUIDs...)
			if err != nil {
				return errors.Trace(err)
			}
			addChanges(processed...)
		case out <- changes.Values():
			changes = nil
			out = nil
		}
	}
}

// Changes returns the channel of obsolete secret changes.
func (w *obsoleteRevisionWatcher) Changes() <-chan []string {
	return w.out
}

// Stop stops the watcher.
func (w *obsoleteRevisionWatcher) Stop() error {
	w.Kill()
	return w.Wait()
}

// Kill kills the watcher via its tomb.
func (w *obsoleteRevisionWatcher) Kill() {
	w.catacomb.Kill(nil)
}

// Wait waits for the watcher's tomb to die,
// and returns the error with which it was killed.
func (w *obsoleteRevisionWatcher) Wait() error {
	return w.catacomb.Wait()
}

// TODO(secrets) - replace with real watcher
func newMockTriggerWatcher(ch watcher.SecretTriggerChannel) *mockSecretTriggerWatcher {
	w := &mockSecretTriggerWatcher{ch: ch}
	w.tomb.Go(func() error {
		<-w.tomb.Dying()
		return tomb.ErrDying
	})
	return w
}

type mockSecretTriggerWatcher struct {
	tomb tomb.Tomb
	ch   watcher.SecretTriggerChannel
}

func (w *mockSecretTriggerWatcher) Changes() watcher.SecretTriggerChannel {
	return w.ch
}

func (w *mockSecretTriggerWatcher) Stop() error {
	w.Kill()
	return w.Wait()
}

func (w *mockSecretTriggerWatcher) Kill() {
	w.tomb.Kill(nil)
}

func (w *mockSecretTriggerWatcher) Err() error {
	return w.tomb.Err()
}

func (w *mockSecretTriggerWatcher) Wait() error {
	return w.tomb.Wait()
}

func (s *WatchableService) WatchSecretRevisionsExpiryChanges(ctx context.Context, owners ...CharmSecretOwner) (watcher.SecretTriggerWatcher, error) {
	ch := make(chan []watcher.SecretTriggerChange, 1)
	ch <- []watcher.SecretTriggerChange{}
	return newMockTriggerWatcher(ch), nil
}

func (s *WatchableService) WatchSecretsRotationChanges(ctx context.Context, owners ...CharmSecretOwner) (watcher.SecretTriggerWatcher, error) {
	ch := make(chan []watcher.SecretTriggerChange, 1)
	ch <- []watcher.SecretTriggerChange{}
	return newMockTriggerWatcher(ch), nil
}

func (s *WatchableService) WatchObsoleteUserSecrets(ctx context.Context) (watcher.NotifyWatcher, error) {
	ch := make(chan struct{}, 1)
	ch <- struct{}{}
	return watchertest.NewMockNotifyWatcher(ch), nil
}

func (s *WatchableService) SecretRotated(
	ctx context.Context, uri *secrets.URI, originalRev int, skip bool, accessor SecretAccessor, token leadership.Token,
) error {
	if err := s.canManage(ctx, uri, accessor, token); err != nil {
		return errors.Trace(err)
	}

	md, err := s.GetSecret(ctx, uri)
	if err != nil {
		return errors.Trace(err)
	}
	if !md.RotatePolicy.WillRotate() {
		s.logger.Debugf("secret %q was rotated but now is set to not rotate")
		return nil
	}
	lastRotateTime := md.NextRotateTime
	if lastRotateTime == nil {
		now := s.clock.Now()
		lastRotateTime = &now
	}
	nextRotateTime := *md.RotatePolicy.NextRotateTime(*lastRotateTime)
	s.logger.Debugf("secret %q was rotated: rev was %d, now %d", uri.ID, originalRev, md.LatestRevision)
	// If the secret will expire before it is due to be next rotated, rotate sooner to allow
	// the charm a chance to update it before it expires.
	willExpire := md.LatestExpireTime != nil && md.LatestExpireTime.Before(nextRotateTime)
	forcedRotateTime := lastRotateTime.Add(secrets.RotateRetryDelay)
	if willExpire {
		s.logger.Warningf("secret %q rev %d will expire before next scheduled rotation", uri.ID, md.LatestRevision)
	}
	if willExpire && forcedRotateTime.Before(*md.LatestExpireTime) || !skip && md.LatestRevision == originalRev {
		nextRotateTime = forcedRotateTime
	}
	s.logger.Debugf("secret %q next rotate time is now: %s", uri.ID, nextRotateTime.UTC().Format(time.RFC3339))

	// TODO(secrets)
	return nil
}
