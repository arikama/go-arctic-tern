package arctictern_test

import (
	"testing"

	"github.com/arikama/go-arctic-tern/arctictern"
	"github.com/arikama/go-mysql-test-container/mysqltestcontainer"
)

func TestSeed(t *testing.T) {
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
	seedDir := "./../seed/example"
	err = arctictern.Seed(db, seedDir)
	if err != nil {
		panic(err)
	}
	err = arctictern.Seed(db, seedDir)
	if err != nil {
		panic(err)
	}
}
