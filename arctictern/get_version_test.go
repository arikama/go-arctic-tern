package arctictern_test

import (
	"fmt"
	"testing"

	"github.com/arikama/go-arctic-tern/arctictern"
	"github.com/stretchr/testify/assert"
)

func TestGetVersion(t *testing.T) {
	expected := 1337
	v := arctictern.GetVersion(fmt.Sprintf("/Users/1/workspace/mysql-test-container/integrationtest/example/resources/db/migration/V%v__create_table.sql", expected))
	assert.Equal(t, expected, v)
}

func TestGetVersionLowerCase(t *testing.T) {
	expected := 1337
	v := arctictern.GetVersion(fmt.Sprintf("/Users/1/workspace/mysql-test-container/integrationtest/example/resources/db/migration/v%v__create_table.sql", expected))
	assert.Equal(t, expected, v)
}

func TestGetVersionMissing(t *testing.T) {
	v := arctictern.GetVersion("abc.sql")
	assert.Equal(t, 0, v)
}

func TestGetVersionMissingVersion(t *testing.T) {
	v := arctictern.GetVersion("V.sql")
	assert.Equal(t, 0, v)
}
