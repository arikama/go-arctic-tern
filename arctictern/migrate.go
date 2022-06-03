package arctictern

import (
	"database/sql"
	"fmt"

	"github.com/arikama/go-mysql-test-container/util"
)

func Migrate(db *sql.DB, migrationDir string) error {
	fmt.Println("🐦 Running migration...")
	_, err := db.Exec(
		`
		CREATE TABLE IF NOT EXISTS migration (
			id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
			version VARCHAR(64) NOT NULL,
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
		);
		`,
	)
	if err != nil {
		return err
	}
	files, err := util.GetFiles(migrationDir)
	if err != nil {
		return err
	}
	for _, file := range files {
		version := fmt.Sprintf("V%v", util.GetVersion(file))
		rows, err := db.Query(
			`
			SELECT id, version
			FROM migration
			WHERE version = ?
			;
			`,
			version,
		)
		if err != nil {
			return err
		}
		if rows.Next() {
			fmt.Printf("🐦 Running migration file %v: skipped\n", version)
			continue
		}
		content, _ := util.LoadFile(file)
		fmt.Printf("🐦 Running migration file %v: %v\n", version, file)
		_, err = db.Exec(content)
		if err != nil {
			return err
		}
		_, err = db.Exec(
			`
			INSERT INTO migration (version)
			VALUES (?)
			;
			`,
			version,
		)
		if err != nil {
			return err
		}
	}
	return nil
}
