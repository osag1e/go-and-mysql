package migrations

import (
	"database/sql"
	"os"
)

func ApplyMigrations(db *sql.DB) error {
	migrationFiles := []string{
		"db/scripts/02_create_books_table.up.sql",
	}

	for _, file := range migrationFiles {
		migrationSQL, err := os.ReadFile(file)
		if err != nil {
			return err
		}

		_, err = db.Exec(string(migrationSQL))
		if err != nil {
			return err
		}
	}

	return nil
}

func RollbackMigrations(db *sql.DB) error {
	migrationFiles := []string{
		"db/scripts/02_create_books_table.down.sql",
	}

	for _, file := range migrationFiles {
		migrationSQL, err := os.ReadFile(file)
		if err != nil {
			return err
		}

		_, err = db.Exec(string(migrationSQL))
		if err != nil {
			return err
		}
	}

	return nil
}
