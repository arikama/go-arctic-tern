package arctictern

import (
	"testing"

	"github.com/arikama/go-mysql-test-container/mysqltestcontainer"
)

func TestMigrate(t *testing.T) {
	db, err := mysqltestcontainer.Start("test", "")
	if err != nil {
		panic(err)
	}
	migrationDir := "./../migration/example"
	err = Migrate(db, migrationDir)
	if err != nil {
		panic(err)
	}
	err = Migrate(db, migrationDir)
	if err != nil {
		panic(err)
	}
}

func TestMigrateInvalid(t *testing.T) {
	db, err := mysqltestcontainer.Start("test", "")
	if err != nil {
		panic(err)
	}
	migrationDir := "./../migration/invalid"
	err = Migrate(db, migrationDir)
	if err != nil && err.Error() != "Error 1064: You have an error in your SQL syntax; check the manual that corresponds to your MySQL server version for the right syntax to use near 'INVALID' at line 1" {
		panic(err)
	}
}
