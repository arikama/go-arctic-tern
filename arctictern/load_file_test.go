package arctictern_test

import (
	"testing"

	"github.com/arikama/go-arctic-tern/arctictern"
)

func TestLoadFile(t *testing.T) {
	content, err := arctictern.LoadFile("./../migration/example/V02.sql")
	if err != nil {
		panic(err)
	}
	expected := "ALTER TABLE `user` ADD COLUMN `username` VARCHAR(64) NOT NULL UNIQUE AFTER `id`;"
	if content != expected {
		t.Errorf("Got=%v, want=%v\n", content, expected)
	}
}

func TestLoadFileMissing(t *testing.T) {
	content, err := arctictern.LoadFile("./../migration/example/missing.sql")
	if err != nil && err.Error() != "open ./../migration/example/missing.sql: no such file or directory" {
		panic(err)
	}
	expected := ""
	if content != expected {
		t.Errorf("Got=%v, want=%v\n", content, expected)
	}
}
