package arctictern_test

import (
	"testing"

	"github.com/arikama/go-arctic-tern/arctictern"
	"github.com/arikama/go-mysql-test-container/mysqltestcontainer"
	"github.com/stretchr/testify/assert"
)

func TestSeed(t *testing.T) {
	result, err := mysqltestcontainer.Start("test")
	assert.Nil(t, err)
	assert.NotNil(t, result)

	db := result.Db
	migrationDir := "./../migration/example"
	err = arctictern.Migrate(db, migrationDir)
	assert.Nil(t, err)

	err = arctictern.Migrate(db, migrationDir)
	assert.Nil(t, err)

	seedDir := "./../seed/example"
	err = arctictern.Seed(db, seedDir)
	assert.Nil(t, err)

	err = arctictern.Seed(db, seedDir)
	assert.Nil(t, err)

	rows, err := db.Query(`SELECT * FROM user;`)
	assert.Nil(t, err)

	count := 0
	for rows.Next() {
		count += 1
	}
	assert.Equal(t, 3, count)
}
