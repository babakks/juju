// Copyright 2024 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package secretsmanager

import (
	"context"

	"github.com/juju/juju/core/leadership"
	"github.com/juju/juju/core/secrets"
	"github.com/juju/juju/core/watcher"
	secretservice "github.com/juju/juju/domain/secret/service"
)

// SecretTriggers instances provide secret rotation/expiry apis.
type SecretTriggers interface {
	WatchSecretRevisionsExpiryChanges(ctx context.Context, owners ...secretservice.CharmSecretOwner) (watcher.SecretTriggerWatcher, error)
	WatchSecretsRotationChanges(ctx context.Context, owners ...secretservice.CharmSecretOwner) (watcher.SecretTriggerWatcher, error)
	WatchObsolete(ctx context.Context, owners ...secretservice.CharmSecretOwner) (watcher.StringsWatcher, error)
	SecretRotated(ctx context.Context, uri *secrets.URI, originalRev int, skip bool, accessor secretservice.SecretAccessor, token leadership.Token) error
}

// SecretsConsumer instances provide secret consumer apis.
type SecretsConsumer interface {
	GetSecretConsumer(ctx context.Context, uri *secrets.URI, unitName string) (*secrets.SecretConsumerMetadata, error)
	GetSecretConsumerAndLatest(ctx context.Context, uri *secrets.URI, unitName string) (*secrets.SecretConsumerMetadata, int, error)
	GetURIByConsumerLabel(ctx context.Context, label string, unitName string) (*secrets.URI, error)
	SaveSecretConsumer(ctx context.Context, uri *secrets.URI, unitName string, md *secrets.SecretConsumerMetadata) error
	GetConsumedRevision(
		ctx context.Context, uri *secrets.URI, unitName string,
		refresh, peek bool, labelToUpdate *string) (int, error)
	WatchConsumedSecretsChanges(ctx context.Context, unitName string) (watcher.StringsWatcher, error)
	GrantSecretAccess(context.Context, *secrets.URI, secretservice.SecretAccessParams) error
	RevokeSecretAccess(context.Context, *secrets.URI, secretservice.SecretAccessParams) error
}

// SecretService provides core secrets operations.
type SecretService interface {
	CreateSecretURIs(ctx context.Context, count int) ([]*secrets.URI, error)
	CreateSecret(context.Context, *secrets.URI, secretservice.CreateSecretParams) error
	UpdateSecret(context.Context, *secrets.URI, secretservice.UpdateSecretParams) error
	DeleteSecret(context.Context, *secrets.URI, secretservice.DeleteSecretParams) error
	GetSecretValue(context.Context, *secrets.URI, int, secretservice.SecretAccessor) (secrets.SecretValue, *secrets.ValueRef, error)
	ListCharmSecrets(context.Context, ...secretservice.CharmSecretOwner) ([]*secrets.SecretMetadata, [][]*secrets.SecretRevisionMetadata, error)
	ProcessSecretConsumerLabel(
		ctx context.Context, unitName string, uri *secrets.URI, label string, token leadership.Token,
	) (*secrets.URI, *string, error)
	ChangeSecretBackend(ctx context.Context, uri *secrets.URI, revision int, params secretservice.ChangeSecretBackendParams) error
	GetSecretGrants(ctx context.Context, uri *secrets.URI, role secrets.SecretRole) ([]secretservice.SecretAccess, error)
}
