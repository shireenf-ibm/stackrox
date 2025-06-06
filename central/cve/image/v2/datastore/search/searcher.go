package search

import (
	"context"

	pgStore "github.com/stackrox/rox/central/cve/image/v2/datastore/store/postgres"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres/schema"
	"github.com/stackrox/rox/pkg/search"
	pkgPostgres "github.com/stackrox/rox/pkg/search/scoped/postgres"
	"github.com/stackrox/rox/pkg/search/sortfields"
)

// Searcher provides search functionality on existing cves.
//
//go:generate mockgen-wrapper
type Searcher interface {
	Search(ctx context.Context, query *v1.Query) ([]search.Result, error)
	Count(ctx context.Context, query *v1.Query) (int, error)
	SearchImageCVEs(context.Context, *v1.Query) ([]*v1.SearchResult, error)
	SearchRawImageCVEs(ctx context.Context, query *v1.Query) ([]*storage.ImageCVEV2, error)
}

// New returns a new instance of Searcher for the given the storage.
func New(storage pgStore.Store) Searcher {
	return &searcherImpl{
		storage:  storage,
		searcher: formatSearcherV2(storage),
	}
}

func formatSearcherV2(searcher search.Searcher) search.Searcher {
	scopedSearcher := pkgPostgres.WithScoping(searcher)
	return sortfields.TransformSortFields(scopedSearcher, schema.ImagesSchema.OptionsMap)
}
