package arctictern

import (
	"database/sql"
	"fmt"

	"github.com/hooligram/kifu"
)

func Migrate(db *sql.DB, migrationDir string) error {
	kifu.Info("Running migration: dir=%v", migrationDir)
	db.Exec(
		`
		CREATE TABLE IF NOT EXISTS migration (
			id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
			version VARCHAR(64) NOT NULL,
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
		);
		`,
	)
	files, err := GetFiles(migrationDir)
	if err != nil {
		return err
	}
	skipped := 0
	for _, file := range files {
		version := fmt.Sprintf("V%v", GetVersion(file))
		rows, _ := db.Query(
			`
			SELECT id, version
			FROM migration
			WHERE version = ?
			;
			`,
			version,
		)
		if rows.Next() {
			skipped += 1
			continue
		}
		content, _ := LoadFile(file)
		statements := GetStatements(content)
		for _, statement := range statements {
			_, err := db.Exec(statement)
			if err != nil {
				return err
			}
		}
		db.Exec(
			`
			INSERT INTO migration (version)
			VALUES (?)
			;
			`,
			version,
		)
	}
	kifu.Info("Migration done: files=%v, skipped=%v", len(files), skipped)
	return nil
}
