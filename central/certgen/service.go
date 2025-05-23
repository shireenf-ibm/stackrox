package certgen

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	clusterDataStore "github.com/stackrox/rox/central/cluster/datastore"
	siStore "github.com/stackrox/rox/central/serviceidentities/datastore"
	"github.com/stackrox/rox/pkg/auth/permissions"
	pkgGRPC "github.com/stackrox/rox/pkg/grpc"
	"github.com/stackrox/rox/pkg/grpc/authz/user"
	"github.com/stackrox/rox/pkg/grpc/routes"
	"github.com/stackrox/rox/pkg/sac/resources"
	"google.golang.org/grpc"
)

// Service represents the certgen service.
type Service interface {
	pkgGRPC.APIServiceWithCustomRoutes
}

type serviceImpl struct {
	clusters          clusterDataStore.DataStore
	serviceIdentities siStore.DataStore
}

func (s *serviceImpl) RegisterServiceServer(_ *grpc.Server) {
}

func (s *serviceImpl) RegisterServiceHandler(_ context.Context, _ *runtime.ServeMux, _ *grpc.ClientConn) error {
	return nil
}

func (s *serviceImpl) CustomRoutes() []routes.CustomRoute {
	return []routes.CustomRoute{
		{
			Route:         "/api/extensions/certgen/central",
			Authorizer:    user.With(permissions.Modify(resources.Administration)),
			ServerHandler: routes.NotImplementedOnManagedServices(http.HandlerFunc(s.centralHandler)),
			Compression:   false,
		},
		{
			Route:         "/api/extensions/certgen/central-db",
			Authorizer:    user.With(permissions.Modify(resources.Administration)),
			ServerHandler: routes.NotImplementedOnManagedServices(http.HandlerFunc(s.centralDBHandler)),
			Compression:   false,
		},
		{
			Route:         "/api/extensions/certgen/scanner",
			Authorizer:    user.With(permissions.Modify(resources.Administration)),
			ServerHandler: routes.NotImplementedOnManagedServices(http.HandlerFunc(s.scannerHandler)),
			Compression:   false,
		},

		{
			Route:         "/api/extensions/certgen/cluster",
			Authorizer:    user.With(permissions.Modify(resources.Administration)),
			ServerHandler: http.HandlerFunc(s.securedClusterHandler),
			Compression:   false,
		},
	}
}

// NewService returns a new certgen service.
func NewService(clusters clusterDataStore.DataStore, serviceIdentities siStore.DataStore) Service {
	return &serviceImpl{
		clusters:          clusters,
		serviceIdentities: serviceIdentities,
	}
}
