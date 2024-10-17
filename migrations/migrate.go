// migrations/migrate.go

package migrate

import (
	"log"

	migrate "github.com/rubenv/sql-migrate"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get SQL DB: %v", err)
	}

	migrations := &migrate.FileMigrationSource{
		Dir: "migrations/sql",
	}

	n, err := migrate.Exec(sqlDB, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatalf("Failed to apply migrations: %v", err)
	}
	log.Printf("Applied %d migrations!", n)
}
