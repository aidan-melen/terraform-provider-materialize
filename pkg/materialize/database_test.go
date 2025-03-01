package materialize

import (
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/MaterializeInc/terraform-provider-materialize/pkg/testhelpers"
	"github.com/jmoiron/sqlx"
)

func TestDatabaseCreate(t *testing.T) {
	testhelpers.WithMockDb(t, func(db *sqlx.DB, mock sqlmock.Sqlmock) {
		mock.ExpectExec(`CREATE DATABASE "database";`).WillReturnResult(sqlmock.NewResult(1, 1))

		if err := NewDatabaseBuilder(db, "database").Create(); err != nil {
			t.Fatal(err)
		}
	})
}
