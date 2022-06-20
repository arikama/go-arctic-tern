package arctictern_test

import (
	"testing"

	"github.com/arikama/go-arctic-tern/arctictern"
	"github.com/stretchr/testify/assert"
)

func TestGetFiles(t *testing.T) {
	files, err := arctictern.GetFiles("./../migration/example")
	assert.Nil(t, err)
	assert.Equal(t, 2, len(files))
}

func TestGetFilesMissing(t *testing.T) {
	_, err := arctictern.GetFiles("./../migration/missing")
	assert.NotNil(t, err)

	if err != nil {
		expected := "open ./../migration/missing: no such file or directory"
		assert.Equal(t, expected, err.Error())
	}
}
