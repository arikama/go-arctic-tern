package arctictern_test

import (
	"testing"

	"github.com/arikama/go-arctic-tern/arctictern"
	"github.com/arikama/go-mysql-test-container/mysqltestcontainer"
)

func TestMigrate(t *testing.T) {
	result, err := mysqltestcontainer.Start("test")
	if err != nil {
		panic(err)
	}
	db := result.Db
	migrationDir := "./../migration/example"
	err = arctictern.Migrate(db, migrationDir)
	if err != nil {
		panic(err)
	}
	err = arctictern.Migrate(db, migrationDir)
	if err != nil {
		panic(err)
	}
}

func TestMigrateInvalid(t *testing.T) {
	result, err := mysqltestcontainer.Start("test")
	if err != nil {
		panic(err)
	}
	db := result.Db
	migrationDir := "./../migration/invalid"
	err = arctictern.Migrate(db, migrationDir)
	if err != nil && err.Error() != "Error 1064: You have an error in your SQL syntax; check the manual that corresponds to your MySQL server version for the right syntax to use near 'INVALID' at line 1" {
		panic(err)
	}
}
