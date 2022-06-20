package arctictern

import (
	"database/sql"
	"fmt"

	"github.com/hooligram/kifu"
)

func Seed(db *sql.DB, seedDir string) error {
	kifu.Info("Running seed: dir=%v", seedDir)
	db.Exec(
		`
		CREATE TABLE IF NOT EXISTS seed (
			id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
			version VARCHAR(64) NOT NULL,
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
		);
		`,
	)
	files, err := GetFiles(seedDir)
	if err != nil {
		panic(err.Error())
	}
	skipped := 0
	for _, file := range files {
		version := fmt.Sprintf("V%v", GetVersion(file))
		rows, _ := db.Query(
			`
			SELECT id, version
			FROM seed
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
			INSERT INTO seed (version)
			VALUES (?)
			;
			`,
			version,
		)
	}
	kifu.Info("Seed done: files=%v, skipped=%v", len(files), skipped)
	return nil
}
