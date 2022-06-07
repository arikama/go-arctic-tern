package arctictern_test

import (
	"fmt"
	"testing"

	"github.com/arikama/go-arctic-tern/arctictern"
)

func TestGetVersion(t *testing.T) {
	expected := 1337
	v := arctictern.GetVersion(fmt.Sprintf("/Users/1/workspace/mysql-test-container/integrationtest/example/resources/db/migration/V%v__create_table.sql", expected))
	if v != expected {
		t.Errorf("Got %v, expected %v\n", v, expected)
	}
}

func TestGetVersionLowerCase(t *testing.T) {
	expected := 1337
	v := arctictern.GetVersion(fmt.Sprintf("/Users/1/workspace/mysql-test-container/integrationtest/example/resources/db/migration/v%v__create_table.sql", expected))
	if v != expected {
		t.Errorf("Got %v, expected %v\n", v, expected)
	}
}

func TestGetVersionMissing(t *testing.T) {
	v := arctictern.GetVersion("abc.sql")
	if v != 0 {
		t.Errorf("Got %v, expected %v\n", v, 1)
	}
}

func TestGetVersionMissingVersion(t *testing.T) {
	v := arctictern.GetVersion("V.sql")
	if v != 0 {
		t.Errorf("Got %v, expected %v\n", v, 1)
	}
}
