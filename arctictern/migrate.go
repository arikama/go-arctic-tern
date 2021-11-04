package arctictern

import (
	"database/sql"
	"fmt"

	"github.com/arikama/go-mysql-test-container/util"
)

func Migrate(db *sql.DB, migrationDir string) error {
	fmt.Println("🐦 Running migration...")
	db.Exec(
		`
		CREATE TABLE IF NOT EXISTS migration (
			id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
			version VARCHAR(64) NOT NULL,
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
		);
		`,
	)
	files, _ := util.GetFiles(migrationDir)
	for _, file := range files {
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
			fmt.Printf("🐦 Running migration file %v: skipped\n", version)
			continue
		}
		content, _ := util.LoadFile(file)
		fmt.Printf("🐦 Running migration file %v: %v\n", version, file)
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
