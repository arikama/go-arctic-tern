package arctictern

import (
	"database/sql"
	"fmt"

	"github.com/arikama/go-mysql-test-container/util"
	"github.com/hooligram/kifu"
)

func Migrate(db *sql.DB, migrationDir string) error {
	kifu.Info("Running migration...")
	db.Exec(
		`
		CREATE TABLE IF NOT EXISTS migration (
			id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
			version VARCHAR(64) NOT NULL,
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
		);
		`,
	)
	files, err := util.GetFiles(migrationDir)
	if err != nil {
		panic(err.Error())
	}
	kifu.Info("Found %v files in %v", len(files), migrationDir)
	for i, file := range files {
		version := fmt.Sprintf("V%v", util.GetVersion(file))
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
			kifu.Info("Running migration file %v: skipped", version)
			continue
		}
		content, _ := util.LoadFile(file)
		kifu.Info("Running migration file #%v: %v", i+1, file)
		_, err := db.Exec(content)
		if err != nil {
			return err
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
	return nil
}
