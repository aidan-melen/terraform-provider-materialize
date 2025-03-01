package materialize

import (
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/MaterializeInc/terraform-provider-materialize/pkg/testhelpers"
	"github.com/jmoiron/sqlx"
)

func TestRoleCreate(t *testing.T) {
	testhelpers.WithMockDb(t, func(db *sqlx.DB, mock sqlmock.Sqlmock) {
		mock.ExpectExec(
			`CREATE ROLE "role" INHERIT CREATEROLE CREATEDB CREATECLUSTER;`,
		).WillReturnResult(sqlmock.NewResult(1, 1))
		b := NewRoleBuilder(db, "role")
		b.Inherit()
		b.CreateRole()
		b.CreateDb()
		b.CreateCluster()

		if err := b.Create(); err != nil {
			t.Fatal(err)
		}
	})
}

func TestRoleAlter(t *testing.T) {
	testhelpers.WithMockDb(t, func(db *sqlx.DB, mock sqlmock.Sqlmock) {
		mock.ExpectExec(
			`ALTER ROLE "role" CREATECLUSTER;`,
		).WillReturnResult(sqlmock.NewResult(1, 1))

		if err := NewRoleBuilder(db, "role").Alter("CREATECLUSTER"); err != nil {
			t.Fatal(err)
		}
	})
}

func TestRoleDrop(t *testing.T) {
	testhelpers.WithMockDb(t, func(db *sqlx.DB, mock sqlmock.Sqlmock) {
		mock.ExpectExec(
			`DROP ROLE "role";`,
		).WillReturnResult(sqlmock.NewResult(1, 1))

		if err := NewRoleBuilder(db, "role").Drop(); err != nil {
			t.Fatal(err)
		}
	})
}
