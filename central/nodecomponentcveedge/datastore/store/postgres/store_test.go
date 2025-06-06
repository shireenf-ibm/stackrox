// Code generated by pg-bindings generator. DO NOT EDIT.

//go:build sql_integration

package postgres

import (
	"context"
	"testing"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres/pgtest"
	"github.com/stackrox/rox/pkg/sac"
	"github.com/stackrox/rox/pkg/testutils"
	"github.com/stretchr/testify/suite"
)

type NodeComponentsCvesEdgesStoreSuite struct {
	suite.Suite
	store  Store
	testDB *pgtest.TestPostgres
}

func TestNodeComponentsCvesEdgesStore(t *testing.T) {
	suite.Run(t, new(NodeComponentsCvesEdgesStoreSuite))
}

func (s *NodeComponentsCvesEdgesStoreSuite) SetupSuite() {

	s.testDB = pgtest.ForT(s.T())
	s.store = New(s.testDB.DB)
}

func (s *NodeComponentsCvesEdgesStoreSuite) SetupTest() {
	ctx := sac.WithAllAccess(context.Background())
	tag, err := s.testDB.Exec(ctx, "TRUNCATE node_components_cves_edges CASCADE")
	s.T().Log("node_components_cves_edges", tag)
	s.store = New(s.testDB.DB)
	s.NoError(err)
}

func (s *NodeComponentsCvesEdgesStoreSuite) TestStore() {
	ctx := sac.WithAllAccess(context.Background())

	store := s.store

	nodeComponentCVEEdge := &storage.NodeComponentCVEEdge{}
	s.NoError(testutils.FullInit(nodeComponentCVEEdge, testutils.SimpleInitializer(), testutils.JSONFieldsFilter))

	foundNodeComponentCVEEdge, exists, err := store.Get(ctx, nodeComponentCVEEdge.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundNodeComponentCVEEdge)

}
