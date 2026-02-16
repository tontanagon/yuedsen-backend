package main

import (
	"fmt"
	"log"


	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
    // Database connection string
    // The user's main.go has "postgresql://". 
    // We try "postgres://" here because the driver registers "postgres" and "postgresql".
    dsn := "postgres://postgres:Yuedsen-1234@db.ordgvmwtiotwbahcyrbu.supabase.co:5432/postgres?sslmode=disable"
    
    // Path to migrations folder relative to current working directory
    sourceURL := "file://db/migrations"

    fmt.Printf("Migrating using Source: %s\n", sourceURL)
	m, err := migrate.New(
		sourceURL,
		dsn,
	)
	if err != nil {
		log.Fatalf("Migration initialization failed: %v", err)
	}

	if err := m.Up(); err != nil {
        if err == migrate.ErrNoChange {
            log.Println("No changes to apply.")
        } else {
		    log.Fatalf("Migration up failed: %v", err)
        }
	} else {
        log.Println("Migration up successful!")
    }
}
