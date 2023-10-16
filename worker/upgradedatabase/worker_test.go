// Copyright 2023 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package upgradedatabase

import (
	"time"

	"github.com/juju/errors"
	names "github.com/juju/names/v4"
	jc "github.com/juju/testing/checkers"
	"github.com/juju/utils/v3"
	"github.com/juju/version/v2"
	"github.com/juju/worker/v3"
	"github.com/juju/worker/v3/dependency"
	"github.com/juju/worker/v3/workertest"
	"go.uber.org/mock/gomock"
	gc "gopkg.in/check.v1"

	coredatabase "github.com/juju/juju/core/database"
	"github.com/juju/juju/core/testing"
	upgrade "github.com/juju/juju/core/upgrade"
	"github.com/juju/juju/core/watcher/watchertest"
	model "github.com/juju/juju/domain/model"
	domainupgrade "github.com/juju/juju/domain/upgrade"
	upgradeerrors "github.com/juju/juju/domain/upgrade/errors"
	databasetesting "github.com/juju/juju/internal/database/testing"
	jujuversion "github.com/juju/juju/version"
)

// baseSuite is embedded in both the worker and manifold tests.
// Tests should not go on this suite directly.

type workerSuite struct {
	baseSuite
	databasetesting.DqliteSuite

	upgradeUUID domainupgrade.UUID
}

var _ = gc.Suite(&workerSuite{})

func (s *workerSuite) TestNewLockSameVersionUnlocked(c *gc.C) {
	defer s.setupMocks(c).Finish()

	s.agentConfig.EXPECT().UpgradedToVersion().Return(jujuversion.Current)
	c.Assert(NewLock(s.agentConfig).IsUnlocked(), jc.IsTrue)
}

func (s *workerSuite) TestNewLockOldVersionLocked(c *gc.C) {
	defer s.setupMocks(c).Finish()

	s.agentConfig.EXPECT().UpgradedToVersion().Return(version.Number{})
	c.Assert(NewLock(s.agentConfig).IsUnlocked(), jc.IsFalse)
}

func (s *workerSuite) TestLockAlreadyUnlocked(c *gc.C) {
	defer s.setupMocks(c).Finish()

	s.lock.EXPECT().IsUnlocked().Return(true)

	w, err := NewUpgradeDatabaseWorker(s.getConfig())
	c.Assert(err, jc.ErrorIsNil)

	err = workertest.CheckKill(c, w)
	c.Check(err, jc.ErrorIs, dependency.ErrUninstall)
}

func (s *workerSuite) TestLockIsUnlockedIfMatchingVersions(c *gc.C) {
	defer s.setupMocks(c).Finish()

	s.lock.EXPECT().IsUnlocked().Return(false)
	s.lock.EXPECT().Unlock()

	cfg := s.getConfig()
	cfg.FromVersion = jujuversion.Current
	cfg.ToVersion = jujuversion.Current

	w, err := NewUpgradeDatabaseWorker(cfg)
	c.Assert(err, jc.ErrorIsNil)

	err = workertest.CheckKill(c, w)
	c.Check(err, jc.ErrorIs, dependency.ErrUninstall)
}

func (s *workerSuite) TestWatchUpgradeInsteadOfPerforming(c *gc.C) {
	defer s.setupMocks(c).Finish()

	// Ensure that the update hasn't already happened.
	s.lock.EXPECT().IsUnlocked().Return(false)

	cfg := s.getConfig()

	ch := make(chan struct{})

	watcher := watchertest.NewMockNotifyWatcher(ch)
	defer workertest.DirtyKill(c, watcher)

	// Walk through the upgrade process:
	//  - Create Upgrade, but it's already started.
	//  - Get the active upgrade.
	//  - Watch for the upgrade to complete.

	srv := s.upgradeService.EXPECT()
	srv.CreateUpgrade(gomock.Any(), cfg.FromVersion, cfg.ToVersion).Return(domainupgrade.UUID(""), upgradeerrors.ErrUpgradeAlreadyStarted)
	srv.ActiveUpgrade(gomock.Any()).Return(s.upgradeUUID, nil)
	srv.WatchForUpgradeState(gomock.Any(), s.upgradeUUID, upgrade.DBCompleted).Return(watcher, nil)

	// We expect the lock to be unlocked when the upgrade completes.
	s.lock.EXPECT().Unlock()

	w, err := NewUpgradeDatabaseWorker(cfg)
	c.Assert(err, jc.ErrorIsNil)

	select {
	case ch <- struct{}{}:
	case <-time.After(testing.ShortWait):
		c.Fatalf("timed out waiting to enqueue change")
	}

	err = workertest.CheckKill(c, w)
	c.Check(err, jc.ErrorIs, dependency.ErrUninstall)
}

func (s *workerSuite) TestWatchUpgradeError(c *gc.C) {
	defer s.setupMocks(c).Finish()

	// Ensure that the update hasn't already happened.
	s.lock.EXPECT().IsUnlocked().Return(false)

	cfg := s.getConfig()

	ch := make(chan struct{})

	watcher := watchertest.NewMockNotifyWatcher(ch)
	defer workertest.DirtyKill(c, watcher)

	// Walk through the upgrade process:
	//  - Create Upgrade, but it's already started.
	//  - Get the active upgrade, but it doesn't exist.

	srv := s.upgradeService.EXPECT()
	srv.CreateUpgrade(gomock.Any(), cfg.FromVersion, cfg.ToVersion).Return(domainupgrade.UUID(""), upgradeerrors.ErrUpgradeAlreadyStarted)
	srv.ActiveUpgrade(gomock.Any()).Return(s.upgradeUUID, errors.NotFoundf("no upgrade"))

	w, err := NewUpgradeDatabaseWorker(cfg)
	c.Assert(err, jc.ErrorIsNil)

	err = workertest.CheckKill(c, w)
	c.Check(err, jc.ErrorIs, dependency.ErrBounce)
}

func (s *workerSuite) TestUpgradeController(c *gc.C) {
	defer s.setupMocks(c).Finish()

	// Ensure that the update hasn't already happened.
	s.lock.EXPECT().IsUnlocked().Return(false)

	cfg := s.getConfig()

	ch := make(chan struct{})

	watcher := watchertest.NewMockNotifyWatcher(ch)
	defer workertest.CheckKill(c, watcher)

	// Walk through the upgrade process:
	//  - Create Upgrade.
	//  - Set the controller ready for upgrade.
	//  - Wait for the upgrade to be ready. This means, all the controller nodes
	//    are synced and ready to be upgraded.
	//  - Start the upgrade, we're the leader.
	//  - Upgrade the controller db.
	//  - Set the db upgrade complete.
	//  - Unlock the lock.

	s.expectStartUpgrade(cfg.FromVersion, cfg.ToVersion, watcher)

	// Controller upgrade.
	s.expectControllerDBUpgrade()

	// Model upgrade (there are no models).
	s.expectModelList([]model.UUID{})

	s.expectDBCompleted()
	done := s.expectUnlock()

	w, err := NewUpgradeDatabaseWorker(cfg)
	c.Assert(err, jc.ErrorIsNil)
	defer workertest.CheckKill(c, w)

	select {
	case ch <- struct{}{}:
	case <-time.After(testing.ShortWait):
		c.Fatalf("timed out waiting to enqueue change")
	}

	select {
	case <-done:
	case <-time.After(testing.LongWait):
		c.Fatalf("timed out waiting for unlock")
	}
}

func (s *workerSuite) TestUpgradeModels(c *gc.C) {
	defer s.setupMocks(c).Finish()

	// Ensure that the update hasn't already happened.
	s.lock.EXPECT().IsUnlocked().Return(false)

	cfg := s.getConfig()

	ch := make(chan struct{})

	watcher := watchertest.NewMockNotifyWatcher(ch)
	defer workertest.CheckKill(c, watcher)

	// Walk through the upgrade process:
	//  - Create Upgrade.
	//  - Set the controller ready for upgrade.
	//  - Wait for the upgrade to be ready. This means, all the controller nodes
	//    are synced and ready to be upgraded.
	//  - Start the upgrade, we're the leader.
	//  - Upgrade the controller db.
	//  - Upgrade all the model dbs.
	//  - Set the db upgrade complete.
	//  - Unlock the lock.

	s.expectStartUpgrade(cfg.FromVersion, cfg.ToVersion, watcher)

	// Controller upgrade.
	s.expectControllerDBUpgrade()

	// Model upgrade.
	modelUUID := model.UUID(utils.MustNewUUID().String())
	s.expectModelList([]model.UUID{modelUUID})
	s.expectModelDBUpgrade(c, modelUUID)

	s.expectDBCompleted()
	done := s.expectUnlock()

	w, err := NewUpgradeDatabaseWorker(cfg)
	c.Assert(err, jc.ErrorIsNil)
	defer workertest.CheckKill(c, w)

	select {
	case ch <- struct{}{}:
	case <-time.After(testing.ShortWait):
		c.Fatalf("timed out waiting to enqueue change")
	}

	select {
	case <-done:
	case <-time.After(testing.LongWait):
		c.Fatalf("timed out waiting for unlock")
	}
}

func (s *workerSuite) getConfig() Config {
	return Config{
		DBUpgradeCompleteLock: s.lock,
		Agent:                 s.agent,
		Logger:                s.logger,
		UpgradeService:        s.upgradeService,
		ControllerNodeService: s.controllerNodeService,
		ModelManagerService:   s.modelManagerService,
		DBGetter:              s.dbGetter,
		FromVersion:           version.MustParse("3.0.0"),
		ToVersion:             version.MustParse("6.6.6"),
		Tag:                   names.NewMachineTag("0"),
	}
}

func (s *workerSuite) setupMocks(c *gc.C) *gomock.Controller {
	ctrl := s.baseSuite.setupMocks(c)

	s.upgradeUUID = domainupgrade.UUID(utils.MustNewUUID().String())

	return ctrl
}

func (s *workerSuite) expectStartUpgrade(from, to version.Number, watcher worker.Worker) {
	srv := s.upgradeService.EXPECT()
	srv.CreateUpgrade(gomock.Any(), from, to).Return(s.upgradeUUID, nil)
	srv.SetControllerReady(gomock.Any(), s.upgradeUUID, "0").Return(nil)
	srv.WatchForUpgradeReady(gomock.Any(), s.upgradeUUID).Return(watcher, nil)
	srv.StartUpgrade(gomock.Any(), s.upgradeUUID).Return(nil)
}

func (s *workerSuite) expectDBCompleted() {
	srv := s.upgradeService.EXPECT()
	srv.SetDBUpgradeCompleted(gomock.Any(), s.upgradeUUID).Return(nil)
}

func (s *workerSuite) expectControllerDBUpgrade() {
	s.controllerNodeService.EXPECT().DqliteNode(gomock.Any(), "0").Return(uint64(0), "192.168.0.1", nil)
	s.dbGetter.EXPECT().GetDB(coredatabase.ControllerNS).Return(s.TxnRunner(), nil)
}

func (s *workerSuite) expectModelList(models []model.UUID) {
	s.modelManagerService.EXPECT().ModelList(gomock.Any()).Return(models, nil)
}

func (s *workerSuite) expectModelDBUpgrade(c *gc.C, modelUUID model.UUID) {
	txnRunner, _ := s.OpenDB(c)
	s.dbGetter.EXPECT().GetDB(modelUUID.String()).Return(txnRunner, nil)
}

func (s *workerSuite) expectUnlock() chan struct{} {
	done := make(chan struct{})
	s.lock.EXPECT().Unlock().DoAndReturn(func() {
		close(done)
	})
	return done
}
