package datastore

import (
	"context"
	"testing"

	"github.com/stackrox/rox/central/notifier/datastore/internal/store"
	pgStore "github.com/stackrox/rox/central/notifier/datastore/internal/store/postgres"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres"
)

// DataStore provides storage functionality for notifiers.
//
//go:generate mockgen-wrapper
type DataStore interface {
	GetNotifier(ctx context.Context, id string) (*storage.Notifier, bool, error)
	GetScrubbedNotifier(ctx context.Context, id string) (*storage.Notifier, bool, error)
	GetNotifiersFiltered(ctx context.Context, filter func(notifier *storage.Notifier) bool) ([]*storage.Notifier, error)
	ForEachNotifier(ctx context.Context, fn func(obj *storage.Notifier) error) error
	GetManyNotifiers(ctx context.Context, notifierIDs []string) ([]*storage.Notifier, error)
	ForEachScrubbedNotifier(ctx context.Context, fn func(obj *storage.Notifier) error) error
	Exists(ctx context.Context, id string) (bool, error)
	AddNotifier(ctx context.Context, notifier *storage.Notifier) (string, error)
	UpdateNotifier(ctx context.Context, notifier *storage.Notifier) error
	UpsertNotifier(ctx context.Context, notifier *storage.Notifier) (string, error)
	UpsertManyNotifiers(ctx context.Context, notifiers []*storage.Notifier) error
	RemoveNotifier(ctx context.Context, id string) error
}

// New returns a new Store instance
func New(storage store.Store) DataStore {
	return &datastoreImpl{
		storage: storage,
	}
}

// GetTestPostgresDataStore provides a datastore connected to postgres for testing purposes.
func GetTestPostgresDataStore(_ testing.TB, pool postgres.DB) DataStore {
	return New(pgStore.New(pool))
}
