package migrations

import (
	"database/sql"
	"fmt"
)

// ApplyMigrations applique les migrations manuellement (si besoin)
// ApplyMigrations применяет миграции вручную (при необходимости)
func ApplyMigrations(db *sql.DB, queries []string) error {
	for _, q := range queries {
		_, err := db.Exec(q)
		if err != nil {
			return fmt.Errorf("migration failed: %w", err)
		}
	}
	return nil
}
