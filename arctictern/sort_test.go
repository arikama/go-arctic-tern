package arctictern_test

import (
	"sort"
	"testing"

	"github.com/arikama/go-arctic-tern/arctictern"
	"github.com/stretchr/testify/assert"
)

func TestByFileVersion(t *testing.T) {
	files := arctictern.ByFileVersion{
		"V00090_test.sql",
		"V00080_test.sql",
		"V012_test.sql",
		"V11_test.sql",
		"V1_test.sql",
		"V000_test.sql",
	}
	expected := []string{
		"V000_test.sql",
		"V1_test.sql",
		"V11_test.sql",
		"V012_test.sql",
		"V00080_test.sql",
		"V00090_test.sql",
	}
	sort.Sort(files)
	for i := 0; i < len(expected); i++ {
		assert.Equal(t, expected[i], files[i])
	}
}
