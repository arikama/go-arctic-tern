package arctictern_test

import (
	"testing"

	"github.com/arikama/go-arctic-tern/arctictern"
	"github.com/stretchr/testify/assert"
)

func TestLoadFile(t *testing.T) {
	expected := "ALTER TABLE `user` ADD COLUMN `username` VARCHAR(64) NOT NULL UNIQUE AFTER `id`;"
	content, err := arctictern.LoadFile("./../migration/example/V02.sql")
	assert.Nil(t, err)
	assert.Equal(t, expected, content)
}

func TestLoadFileMissing(t *testing.T) {
	content, err := arctictern.LoadFile("./../migration/example/missing.sql")
	assert.NotNil(t, err)
	assert.Equal(t, "", content)
}
