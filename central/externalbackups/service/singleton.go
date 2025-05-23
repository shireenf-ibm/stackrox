package service

import (
	"context"

	"github.com/stackrox/rox/central/externalbackups/datastore"
	"github.com/stackrox/rox/central/externalbackups/manager"
	"github.com/stackrox/rox/central/integrationhealth/reporter"
	backupListener "github.com/stackrox/rox/central/systeminfo/listener"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/sac"
	"github.com/stackrox/rox/pkg/sac/resources"
	"github.com/stackrox/rox/pkg/sync"
)

var (
	once sync.Once

	as Service
)

func initialize() {
	mgr := initializeManager()
	as = New(datastore.Singleton(), reporter.Singleton(), mgr)
}

func initializeManager() manager.Manager {
	ctx := sac.WithGlobalAccessScopeChecker(context.Background(),
		sac.AllowFixedScopes(
			sac.AccessModeScopeKeys(storage.Access_READ_ACCESS, storage.Access_READ_WRITE_ACCESS),
			sac.ResourceScopeKeys(resources.Integration)))

	mgr := manager.New(reporter.Singleton(), backupListener.Singleton())
	err := datastore.Singleton().ForEachBackup(ctx, func(b *storage.ExternalBackup) error {
		if err := mgr.Upsert(ctx, b); err != nil {
			log.Errorf("error initializing backup: %v", err)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return mgr
}

// Singleton provides the instance of the Service interface to register.
func Singleton() Service {
	once.Do(initialize)
	return as
}
