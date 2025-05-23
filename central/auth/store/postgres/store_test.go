// Code generated by pg-bindings generator. DO NOT EDIT.

//go:build sql_integration

package postgres

import (
	"context"
	"testing"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres/pgtest"
	"github.com/stackrox/rox/pkg/protoassert"
	"github.com/stackrox/rox/pkg/sac"
	"github.com/stackrox/rox/pkg/search"
	"github.com/stackrox/rox/pkg/testutils"
	"github.com/stretchr/testify/suite"
)

type AuthMachineToMachineConfigsStoreSuite struct {
	suite.Suite
	store  Store
	testDB *pgtest.TestPostgres
}

func TestAuthMachineToMachineConfigsStore(t *testing.T) {
	suite.Run(t, new(AuthMachineToMachineConfigsStoreSuite))
}

func (s *AuthMachineToMachineConfigsStoreSuite) SetupSuite() {

	s.testDB = pgtest.ForT(s.T())
	s.store = New(s.testDB.DB)
}

func (s *AuthMachineToMachineConfigsStoreSuite) SetupTest() {
	ctx := sac.WithAllAccess(context.Background())
	tag, err := s.testDB.Exec(ctx, "TRUNCATE auth_machine_to_machine_configs CASCADE")
	s.T().Log("auth_machine_to_machine_configs", tag)
	s.store = New(s.testDB.DB)
	s.NoError(err)
}

func (s *AuthMachineToMachineConfigsStoreSuite) TestStore() {
	ctx := sac.WithAllAccess(context.Background())

	store := s.store

	authMachineToMachineConfig := &storage.AuthMachineToMachineConfig{}
	s.NoError(testutils.FullInit(authMachineToMachineConfig, testutils.SimpleInitializer(), testutils.JSONFieldsFilter))

	foundAuthMachineToMachineConfig, exists, err := store.Get(ctx, authMachineToMachineConfig.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundAuthMachineToMachineConfig)

	withNoAccessCtx := sac.WithNoAccess(ctx)

	s.NoError(store.Upsert(ctx, authMachineToMachineConfig))
	foundAuthMachineToMachineConfig, exists, err = store.Get(ctx, authMachineToMachineConfig.GetId())
	s.NoError(err)
	s.True(exists)
	protoassert.Equal(s.T(), authMachineToMachineConfig, foundAuthMachineToMachineConfig)

	authMachineToMachineConfigCount, err := store.Count(ctx, search.EmptyQuery())
	s.NoError(err)
	s.Equal(1, authMachineToMachineConfigCount)
	authMachineToMachineConfigCount, err = store.Count(withNoAccessCtx, search.EmptyQuery())
	s.NoError(err)
	s.Zero(authMachineToMachineConfigCount)

	authMachineToMachineConfigExists, err := store.Exists(ctx, authMachineToMachineConfig.GetId())
	s.NoError(err)
	s.True(authMachineToMachineConfigExists)
	s.NoError(store.Upsert(ctx, authMachineToMachineConfig))
	s.ErrorIs(store.Upsert(withNoAccessCtx, authMachineToMachineConfig), sac.ErrResourceAccessDenied)

	s.NoError(store.Delete(ctx, authMachineToMachineConfig.GetId()))
	foundAuthMachineToMachineConfig, exists, err = store.Get(ctx, authMachineToMachineConfig.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundAuthMachineToMachineConfig)
	s.ErrorIs(store.Delete(withNoAccessCtx, authMachineToMachineConfig.GetId()), sac.ErrResourceAccessDenied)

	var authMachineToMachineConfigs []*storage.AuthMachineToMachineConfig
	var authMachineToMachineConfigIDs []string
	for i := 0; i < 200; i++ {
		authMachineToMachineConfig := &storage.AuthMachineToMachineConfig{}
		s.NoError(testutils.FullInit(authMachineToMachineConfig, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))
		authMachineToMachineConfigs = append(authMachineToMachineConfigs, authMachineToMachineConfig)
		authMachineToMachineConfigIDs = append(authMachineToMachineConfigIDs, authMachineToMachineConfig.GetId())
	}

	s.NoError(store.UpsertMany(ctx, authMachineToMachineConfigs))

	authMachineToMachineConfigCount, err = store.Count(ctx, search.EmptyQuery())
	s.NoError(err)
	s.Equal(200, authMachineToMachineConfigCount)

	s.NoError(store.DeleteMany(ctx, authMachineToMachineConfigIDs))

	authMachineToMachineConfigCount, err = store.Count(ctx, search.EmptyQuery())
	s.NoError(err)
	s.Equal(0, authMachineToMachineConfigCount)
}
