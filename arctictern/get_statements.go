package arctictern

import (
	"fmt"
	"strings"
)

func GetStatements(raw string) []string {
	statements := []string{}
	for _, s := range strings.Split(raw, ";") {
		trimmed := strings.TrimSpace(s)
		if trimmed != "" {
			statements = append(statements, fmt.Sprintf("%v;", trimmed))
		}
	}
	return statements
}
