package arctictern_test

import (
	"testing"

	"github.com/arikama/go-arctic-tern/arctictern"
)

func TestGetFiles(t *testing.T) {
	files, err := arctictern.GetFiles("./../migration/example")
	if err != nil {
		t.Errorf("%v\n", err.Error())
	}
	expected := 2
	if len(files) != expected {
		t.Errorf("Want=%v, got=%v\n", expected, len(files))
	}
}

func TestGetFilesMissing(t *testing.T) {
	_, err := arctictern.GetFiles("./../migration/missing")
	if err != nil {
		expected := "open ./../migration/missing: no such file or directory"
		if err.Error() != expected {
			t.Errorf(expected, err.Error())
		}
	}
}
