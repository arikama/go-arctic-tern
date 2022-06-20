package arctictern_test

import (
	"testing"

	"github.com/arikama/go-arctic-tern/arctictern"
	"github.com/stretchr/testify/assert"
)

func TestGetStatements(t *testing.T) {
	s := `
		INSERT INTO user (username) VALUES ("awice");
		INSERT INTO user (username) VALUES ("numb3r5");
		INSERT INTO user (username) VALUES ("LarryNY");
	`
	statements := arctictern.GetStatements(s)

	assert.Equal(t, 3, len(statements))
	assert.Equal(t, `INSERT INTO user (username) VALUES ("awice");`, statements[0])
	assert.Equal(t, `INSERT INTO user (username) VALUES ("numb3r5");`, statements[1])
	assert.Equal(t, `INSERT INTO user (username) VALUES ("LarryNY");`, statements[2])
}
