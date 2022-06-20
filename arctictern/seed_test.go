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

func TestSeedInvalid(t *testing.T) {
	result, err := mysqltestcontainer.Start("test")
	assert.Nil(t, err)
	assert.NotNil(t, result)

	err = arctictern.Migrate(result.Db, "./../migration/example")
	assert.Nil(t, err)

	err = arctictern.Seed(result.Db, "./../seed/invalid")
	assert.NotNil(t, err)
	assert.Equal(t,
		`Error 1064: You have an error in your SQL syntax; check the manual that corresponds to your MySQL server version for the right syntax to use near 'INVALID' at line 1`,
		err.Error(),
	)
}

func TestSeedMissing(t *testing.T) {
	result, err := mysqltestcontainer.Start("test")
	assert.Nil(t, err)
	assert.NotNil(t, result)

	err = arctictern.Migrate(result.Db, "./../migration/example")
	assert.Nil(t, err)

	err = arctictern.Seed(result.Db, "./../seed/missing")
	assert.NotNil(t, err)
	assert.Equal(t,
		`open ./../seed/missing: no such file or directory`,
		err.Error(),
	)
}

func TestSeedStress(t *testing.T) {
	result, err := mysqltestcontainer.Start("test")
	assert.Nil(t, err)
	assert.NotNil(t, result)

	err = arctictern.Migrate(result.Db, "./../migration/stress")
	assert.Nil(t, err)

	err = arctictern.Seed(result.Db, "./../seed/stress")
	assert.Nil(t, err)

	rows, err := result.Db.Query("SELECT COUNT(*) FROM user;")
	assert.Nil(t, err)
	rows.Next()
	var count int
	rows.Scan(&count)
	assert.Equal(t, 2000, count)
}
