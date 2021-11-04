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
	migrationDir := "./../migrationexample"
	err = Migrate(db, migrationDir)
	if err != nil {
		panic(err)
	}
	err = Migrate(db, migrationDir)
	if err != nil {
		panic(err)
	}
}
