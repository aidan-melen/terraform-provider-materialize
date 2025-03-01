package resources

import (
	"context"
	"testing"

	"github.com/MaterializeInc/terraform-provider-materialize/pkg/testhelpers"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

var inMaterializedView = map[string]interface{}{
	"name":          "materialized_view",
	"schema_name":   "schema",
	"database_name": "database",
	"statement":     "SELECT 1 FROM 1",
}

var readMaterializedView string = `
SELECT
	mz_materialized_views.id,
	mz_materialized_views.name AS materialized_view_name,
	mz_schemas.name AS schema_name,
	mz_databases.name AS database_name,
	mz_clusters.name AS cluster_name
FROM mz_materialized_views
JOIN mz_schemas
	ON mz_materialized_views.schema_id = mz_schemas.id
JOIN mz_databases
	ON mz_schemas.database_id = mz_databases.id
LEFT JOIN mz_clusters
	ON mz_materialized_views.cluster_id = mz_clusters.id
WHERE mz_materialized_views.id = 'u1';`

func TestResourceMaterializedViewCreate(t *testing.T) {
	r := require.New(t)
	d := schema.TestResourceDataRaw(t, MaterializedView().Schema, inMaterializedView)
	r.NotNil(d)

	testhelpers.WithMockDb(t, func(db *sqlx.DB, mock sqlmock.Sqlmock) {
		// Create
		mock.ExpectExec(`CREATE MATERIALIZED VIEW "database"."schema"."materialized_view" AS SELECT 1 FROM 1;`).WillReturnResult(sqlmock.NewResult(1, 1))

		// Query Id
		ir := mock.NewRows([]string{"id", "materialized_view_name", "schema_name", "database_name", "cluster_name"}).
			AddRow("u1", "view", "schema", "database", "cluster")
		mock.ExpectQuery(`
			SELECT
				mz_materialized_views.id,
				mz_materialized_views.name AS materialized_view_name,
				mz_schemas.name AS schema_name,
				mz_databases.name AS database_name,
				mz_clusters.name AS cluster_name
			FROM mz_materialized_views
			JOIN mz_schemas
				ON mz_materialized_views.schema_id = mz_schemas.id
			JOIN mz_databases
				ON mz_schemas.database_id = mz_databases.id
			LEFT JOIN mz_clusters
				ON mz_materialized_views.cluster_id = mz_clusters.id
			WHERE mz_databases.name = 'database'
			AND mz_materialized_views.name = 'materialized_view'
			AND mz_schemas.name = 'schema';
		`).WillReturnRows(ir)

		// Query Params
		ip := mock.NewRows([]string{"id", "materialized_view_name", "schema_name", "database_name", "cluster_name"}).
			AddRow("u1", "view", "schema", "database", "cluster")
		mock.ExpectQuery(readMaterializedView).WillReturnRows(ip)

		if err := materializedViewCreate(context.TODO(), d, db); err != nil {
			t.Fatal(err)
		}
	})

}

func TestResourceMaterializedViewUpdate(t *testing.T) {
	r := require.New(t)
	d := schema.TestResourceDataRaw(t, MaterializedView().Schema, inMaterializedView)

	// Set current state
	d.SetId("u1")
	d.Set("name", "old_materialized_view")
	r.NotNil(d)

	testhelpers.WithMockDb(t, func(db *sqlx.DB, mock sqlmock.Sqlmock) {
		mock.ExpectExec(`ALTER MATERIALIZED VIEW "database"."schema"."old_materialized_view" RENAME TO "database"."schema"."materialized_view";`).WillReturnResult(sqlmock.NewResult(1, 1))

		// Query Params
		ip := mock.NewRows([]string{"id", "materialized_view_name", "schema_name", "database_name", "cluster_name"}).
			AddRow("u1", "view", "schema", "database", "cluster")
		mock.ExpectQuery(readMaterializedView).WillReturnRows(ip)

		if err := materializedViewUpdate(context.TODO(), d, db); err != nil {
			t.Fatal(err)
		}
	})

}

func TestResourceMaterializedViewDelete(t *testing.T) {
	r := require.New(t)

	in := map[string]interface{}{
		"name":          "materialized_view",
		"schema_name":   "schema",
		"database_name": "database",
	}
	d := schema.TestResourceDataRaw(t, MaterializedView().Schema, in)
	r.NotNil(d)

	testhelpers.WithMockDb(t, func(db *sqlx.DB, mock sqlmock.Sqlmock) {
		mock.ExpectExec(`DROP MATERIALIZED VIEW "database"."schema"."materialized_view";`).WillReturnResult(sqlmock.NewResult(1, 1))

		if err := materializedViewDelete(context.TODO(), d, db); err != nil {
			t.Fatal(err)
		}
	})

}
