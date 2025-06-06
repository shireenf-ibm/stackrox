//go:build sql_integration

package m201tom202

import (
	"context"
	"testing"
	"time"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/migrator/migrations/m_202_to_m_203_vuln_requests_for_suppressed_cves/schema"
	imageCVEEdgeStore "github.com/stackrox/rox/migrator/migrations/m_202_to_m_203_vuln_requests_for_suppressed_cves/store/imagecveedges"
	imageCVEStore "github.com/stackrox/rox/migrator/migrations/m_202_to_m_203_vuln_requests_for_suppressed_cves/store/imagecves"
	imageStore "github.com/stackrox/rox/migrator/migrations/m_202_to_m_203_vuln_requests_for_suppressed_cves/store/images"
	vulnReqStore "github.com/stackrox/rox/migrator/migrations/m_202_to_m_203_vuln_requests_for_suppressed_cves/store/vulnerabilityrequests"
	pghelper "github.com/stackrox/rox/migrator/migrations/postgreshelper"
	"github.com/stackrox/rox/migrator/types"
	"github.com/stackrox/rox/pkg/postgres/pgutils"
	"github.com/stackrox/rox/pkg/protocompat"
	"github.com/stackrox/rox/pkg/protoconv"
	"github.com/stackrox/rox/pkg/sac"
	"github.com/stackrox/rox/pkg/search"
	"github.com/stackrox/rox/pkg/set"
	"github.com/stackrox/rox/pkg/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type migrationTestSuite struct {
	suite.Suite

	db  *pghelper.TestPostgres
	ctx context.Context
}

func TestMigration(t *testing.T) {
	suite.Run(t, new(migrationTestSuite))
}

func (s *migrationTestSuite) SetupSuite() {
	s.ctx = sac.WithAllAccess(context.Background())
	s.db = pghelper.ForT(s.T(), false)

	pgutils.CreateTableFromModel(s.ctx, s.db.GetGormDB(), schema.CreateTableImagesStmt)
	pgutils.CreateTableFromModel(s.ctx, s.db.GetGormDB(), schema.CreateTableImageCvesStmt)
	pgutils.CreateTableFromModel(s.ctx, s.db.GetGormDB(), schema.CreateTableImageCveEdgesStmt)
	pgutils.CreateTableFromModel(s.ctx, s.db.GetGormDB(), schema.CreateTableVulnerabilityRequestsStmt)
}

func (s *migrationTestSuite) TestMigration() {
	images := []*storage.Image{
		getTestImage("image1"),
		getTestImage("image2"),
	}
	cves := []*storage.ImageCVE{
		getTestImageCVE("cve-2023-123", true, false),
		getTestImageCVE("cve-2023-124", true, false),
		getTestImageCVE("cve-2023-125", true, false),
		getTestImageCVE("cve-2023-126", true, true),
		getTestImageCVE("cve-2023-127", true, true),
		getTestImageCVE("cve-2023-128", true, true),
		getTestImageCVE("cve-2023-129", false, false),
		getTestImageCVE("cve-2023-131", false, false),
		getTestImageCVE("cve-2023-132", false, false),
		getTestImageCVE("cve-2023-134", true, false), // has conflicting vuln exception.
		getTestImageCVE("cve-2023-135", true, false), // has conflicting vuln exception.
		getTestImageCVE("cve-2023-136", true, false),
		getTestImageCVE("cve-2023-137", true, false),
		func() *storage.ImageCVE {
			cve := getTestImageCVE("cve-2023-138", true, false)
			cve.SnoozeExpiry = nil
			return cve
		}(),
	}
	unExpiredSnoozedCVEMap := map[string]*storage.ImageCVE{
		"cve-2023-123": cves[0],
		"cve-2023-124": cves[1],
		"cve-2023-125": cves[2],
		"cve-2023-136": cves[11],
		"cve-2023-137": cves[12],
		"cve-2023-138": cves[13],
	}

	imageCVEEdges := []*storage.ImageCVEEdge{
		getTestImageCVEEdge("image1", "cve-2023-123"),
		getTestImageCVEEdge("image1", "cve-2023-124"),
		getTestImageCVEEdge("image1", "cve-2023-125"),
		getTestImageCVEEdge("image1", "cve-2023-129"),
		getTestImageCVEEdge("image1", "cve-2023-131"),
		getTestImageCVEEdge("image1", "cve-2023-132"),
	}

	now := protocompat.TimestampNow()
	expiry := &protocompat.Timestamp{
		Seconds: now.Seconds + int64(7*24*time.Hour.Seconds()),
		Nanos:   now.Nanos,
	}
	exceptions := []*storage.VulnerabilityRequest{
		createVulnerabilityRequest("cve-2023-134", now, expiry),
		createVulnerabilityRequest("cve-2023-135", now, expiry),
		fakeImageScopeVulnReq("registry", "remote", ".*", "cve-2023-136"),
		fakeImageScopeVulnReq("registry", "remote", "tag", "cve-2023-137"),
	}
	existingExceptionsIDs := set.NewStringSet()
	for _, existingException := range exceptions {
		existingExceptionsIDs.Add(existingException.GetId())
	}
	require.NoError(s.T(), imageStore.New(s.db).UpsertMany(s.ctx, images))
	require.NoError(s.T(), imageCVEStore.New(s.db).UpsertMany(s.ctx, cves))
	require.NoError(s.T(), imageCVEEdgeStore.New(s.db).UpsertMany(s.ctx, imageCVEEdges))
	require.NoError(s.T(), vulnReqStore.New(s.db).UpsertMany(s.ctx, exceptions))

	dbs := &types.Databases{
		GormDB:     s.db.GetGormDB(),
		PostgresDB: s.db.DB,
		DBCtx:      s.ctx,
	}

	s.Require().NoError(migration.Run(dbs))

	// Verify exception is created for all unexpired and non-conflicting snoozed vulnerabilities.
	newStore := vulnReqStore.New(s.db)
	objs, err := newStore.GetByQuery(s.ctx, search.EmptyQuery())
	assert.NoError(s.T(), err)
	assertVulnReq(s.T(), existingExceptionsIDs, unExpiredSnoozedCVEMap, objs)

	// Verify search works.
	objs, err = newStore.GetByQuery(s.ctx,
		search.NewQueryBuilder().AddExactMatches(search.CVE, "cve-2023-123").ProtoQuery())
	assert.NoError(s.T(), err)
	assertVulnReq(s.T(), existingExceptionsIDs, unExpiredSnoozedCVEMap, objs)

	// Verify search works.
	objs, err = newStore.GetByQuery(s.ctx,
		search.NewQueryBuilder().
			AddExactMatches(search.RequestedVulnerabilityState, storage.VulnerabilityState_DEFERRED.String()).ProtoQuery())
	assert.NoError(s.T(), err)
	assertVulnReq(s.T(), existingExceptionsIDs, unExpiredSnoozedCVEMap, objs)

	// Verify expired snoozed CVEs do not have exception request.
	objs, err = newStore.GetByQuery(s.ctx,
		search.NewQueryBuilder().AddExactMatches(search.CVE, "cve-2023-127").ProtoQuery())
	assert.NoError(s.T(), err)
	assert.Empty(s.T(), objs)

	// Verify observed CVEs do not have exception request.
	objs, err = newStore.GetByQuery(s.ctx,
		search.NewQueryBuilder().AddExactMatches(search.CVE, "cve-2023-132").ProtoQuery())
	assert.NoError(s.T(), err)
	assert.Empty(s.T(), objs)

	// Verify storage.ImageCVEEdge is updated.
	var edgeObjs []*storage.ImageCVEEdge
	err = imageCVEEdgeStore.New(s.db).Walk(s.ctx, func(obj *storage.ImageCVEEdge) error {
		if obj.State == storage.VulnerabilityState_DEFERRED {
			edgeObjs = append(edgeObjs, obj)
		}
		return nil
	})
	assert.NoError(s.T(), err)
	assert.Len(s.T(), edgeObjs, 3)
	for _, obj := range edgeObjs {
		assert.NotNil(s.T(), unExpiredSnoozedCVEMap[obj.GetImageCveId()])
	}

	// Verify storage.ImageCVE field is not reset.
	cveObjs, err := imageCVEStore.New(s.db).GetByQuery(s.ctx, search.EmptyQuery())
	assert.NoError(s.T(), err)
	assert.Len(s.T(), cveObjs, len(cves))
	for _, obj := range cveObjs {
		for _, cve := range cves {
			if cve.GetCveBaseInfo().GetCve() == obj.GetCveBaseInfo().GetCve() {
				assert.Equal(s.T(), cve.GetSnoozed(), obj.GetSnoozed())
			}
		}
	}
}

func assertVulnReq(
	t *testing.T,
	existingExceptions set.StringSet,
	unExpired map[string]*storage.ImageCVE,
	exceptions []*storage.VulnerabilityRequest,
) {
	for _, obj := range exceptions {
		if existingExceptions.Contains(obj.GetId()) {
			continue
		}

		cve := unExpired[obj.GetCves().GetCves()[0]]
		assert.NotNil(t, cve)
		assert.False(t, obj.GetExpired())
		assert.Equal(t, cve.GetSnoozeExpiry(), obj.GetDeferralReq().GetExpiry().GetExpiresOn())
		assert.Equal(t, ".*", obj.GetScope().GetImageScope().GetRegistry())
		assert.Equal(t, ".*", obj.GetScope().GetImageScope().GetRemote())
		assert.Equal(t, ".*", obj.GetScope().GetImageScope().GetTag())
		assert.NotEmpty(t, obj.GetComments())
		assert.NotEmpty(t, obj.GetApproversV2())
	}
}

func getTestImageCVE(cve string, snoozed, expired bool) *storage.ImageCVE {
	return &storage.ImageCVE{
		Id: cve,
		CveBaseInfo: &storage.CVEInfo{
			Cve: cve,
		},
		Snoozed: snoozed,
		SnoozeExpiry: func() *protocompat.Timestamp {
			now := time.Now()
			if expired {
				now = now.Add(-(7 * 24 * time.Hour))
			} else {
				now = now.Add(7 * 24 * time.Hour)
			}
			return protoconv.ConvertTimeToTimestamp(now)
		}(),
	}
}

func getTestImage(image string) *storage.Image {
	return &storage.Image{
		Id: image,
	}
}

func getTestImageCVEEdge(image, cve string) *storage.ImageCVEEdge {
	return &storage.ImageCVEEdge{
		Id:         uuid.NewV4().String(),
		ImageCveId: cve,
		ImageId:    image,
	}
}

func fakeImageScopeVulnReq(registry, remote, tag string, cve string) *storage.VulnerabilityRequest {
	return &storage.VulnerabilityRequest{
		Id:          uuid.NewV4().String(),
		Name:        cve,
		TargetState: storage.VulnerabilityState_DEFERRED,
		Status:      storage.RequestStatus_APPROVED,
		Expired:     false,
		Scope: &storage.VulnerabilityRequest_Scope{
			Info: &storage.VulnerabilityRequest_Scope_ImageScope{
				ImageScope: &storage.VulnerabilityRequest_Scope_Image{
					Registry: registry,
					Remote:   remote,
					Tag:      tag,
				},
			},
		},
		Entities: &storage.VulnerabilityRequest_Cves{
			Cves: &storage.VulnerabilityRequest_CVEs{
				Cves: []string{cve},
			},
		},
	}
}
