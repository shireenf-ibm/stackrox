package telemetry

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/stackrox/rox/central/authprovider/datastore"
	groupDataStore "github.com/stackrox/rox/central/group/datastore"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/sac"
	"github.com/stackrox/rox/pkg/sac/resources"
	"github.com/stackrox/rox/pkg/set"
	"github.com/stackrox/rox/pkg/telemetry/phonehome"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Gather auth provider names and number of groups per auth provider.
var Gather phonehome.GatherFunc = func(ctx context.Context) (map[string]any, error) {
	ctx = sac.WithGlobalAccessScopeChecker(ctx,
		sac.AllowFixedScopes(
			sac.AccessModeScopeKeys(storage.Access_READ_ACCESS),
			sac.ResourceScopeKeys(resources.Access)))
	props := make(map[string]any)
	providerIDTypes := make(map[string]string)
	providerTypes := set.NewSet[string]()
	providerOriginCount := map[storage.Traits_Origin]int{
		storage.Traits_DEFAULT:              0,
		storage.Traits_IMPERATIVE:           0,
		storage.Traits_DECLARATIVE:          0,
		storage.Traits_DECLARATIVE_ORPHANED: 0,
	}

	err := datastore.Singleton().ForEachAuthProvider(ctx, func(provider *storage.AuthProvider) error {
		providerIDTypes[provider.GetId()] = provider.GetType()
		providerTypes.Add(provider.GetType())
		providerOriginCount[provider.GetTraits().GetOrigin()]++
		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to process auth providers for telemetry")
	}

	props["Auth Providers"] = providerTypes.AsSlice()
	providerGroups := make(map[string]int)

	err = groupDataStore.Singleton().ForEach(ctx, func(group *storage.Group) error {
		id := group.GetProps().GetAuthProviderId()
		providerGroups[id] = providerGroups[id] + 1
		return nil
	})
	if err != nil {
		return props, errors.Wrap(err, "failed to get Groups")
	}

	for id, n := range providerGroups {
		props["Total Groups of "+providerIDTypes[id]] = n
	}

	for origin, count := range providerOriginCount {
		props[fmt.Sprintf("Total %s Auth Providers",
			cases.Title(language.English, cases.Compact).String(origin.String()))] = count
	}
	return props, nil
}
