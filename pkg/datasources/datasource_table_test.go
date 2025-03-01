package datasources

import (
	"context"
	"testing"

	"github.com/MaterializeInc/terraform-provider-materialize/pkg/testhelpers"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

func TestTableDatasource(t *testing.T) {
	r := require.New(t)

	in := map[string]interface{}{
		"schema_name":   "schema",
		"database_name": "database",
	}
	d := schema.TestResourceDataRaw(t, Table().Schema, in)
	r.NotNil(d)

	testhelpers.WithMockDb(t, func(db *sqlx.DB, mock sqlmock.Sqlmock) {
		ir := mock.NewRows([]string{"id", "name", "schema_name", "database_name"}).
			AddRow("u1", "table", "schema", "database")
		mock.ExpectQuery(`
			SELECT
				mz_tables.id,
				mz_tables.name,
				mz_schemas.name AS schema_name,
				mz_databases.name AS database_name
			FROM mz_tables
			JOIN mz_schemas
				ON mz_tables.schema_id = mz_schemas.id
			JOIN mz_databases
				ON mz_schemas.database_id = mz_databases.id
			WHERE mz_databases.name = 'database'
			AND mz_schemas.name = 'schema';`).WillReturnRows(ir)

		if err := tableRead(context.TODO(), d, db); err != nil {
			t.Fatal(err)
		}
	})

}
