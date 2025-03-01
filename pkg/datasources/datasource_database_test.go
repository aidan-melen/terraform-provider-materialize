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

func TestDatabaseDatasource(t *testing.T) {
	r := require.New(t)

	in := map[string]interface{}{}
	d := schema.TestResourceDataRaw(t, Database().Schema, in)
	r.NotNil(d)

	testhelpers.WithMockDb(t, func(db *sqlx.DB, mock sqlmock.Sqlmock) {
		ir := mock.NewRows([]string{"id", "database_name"}).
			AddRow("u1", "database")
		mock.ExpectQuery(`SELECT id, name AS database_name FROM mz_databases;`).WillReturnRows(ir)

		if err := databaseRead(context.TODO(), d, db); err != nil {
			t.Fatal(err)
		}
	})

}
