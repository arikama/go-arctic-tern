package arctictern_test

import (
	"testing"

	"github.com/arikama/go-arctic-tern/arctictern"
	"github.com/arikama/go-mysql-test-container/mysqltestcontainer"
	"github.com/stretchr/testify/assert"
)

func TestMigrate(t *testing.T) {
	result, err := mysqltestcontainer.Start("test")
	assert.Nil(t, err)
	assert.NotNil(t, result)

	db := result.Db
	migrationDir := "./../migration/example"
	err = arctictern.Migrate(db, migrationDir)
	assert.Nil(t, err)

	err = arctictern.Migrate(db, migrationDir)
	assert.Nil(t, err)
}

func TestMigrateInvalid(t *testing.T) {
	result, err := mysqltestcontainer.Start("test")
	assert.Nil(t, err)
	assert.NotNil(t, result)

	db := result.Db
	migrationDir := "./../migration/invalid"
	err = arctictern.Migrate(db, migrationDir)
	assert.NotNil(t, err)
	assert.Equal(t,
		"Error 1064: You have an error in your SQL syntax; check the manual that corresponds to your MySQL server version for the right syntax to use near 'INVALID' at line 1",
		err.Error(),
	)
}
