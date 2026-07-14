package internal

import (
	"context"
	"database/sql"
	"embed"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "modernc.org/sqlite"
)

func RunApp() {

	// Connect to DB
	sqliteDb, err := sql.Open("sqlite", "test.db")
	if err != nil {
		panic(err.Error())
	}

	// Ensure DB present and run migrations
	if err := ensureDB(sqliteDb); err != nil {
		panic(err.Error())
	}

	// Register services, inject dependencies, register routes

	fmt.Println("running")

}

//go:embed migrations/*.sql
var migrations embed.FS

func ensureDB(db *sql.DB) error {
	ctx := context.Background()

	// DB healthcheck
	err := db.PingContext(ctx)
	if err != nil {
		return err
	}

	// Run migrations
	src, err := iofs.New(migrations, "migrations")
	if err != nil {
		return err
	}

	// Create driver to access db
	drv, err := sqlite.WithInstance(db, &sqlite.Config{})
	if err != nil {
		return err
	}

	migrator, err := migrate.NewWithInstance("iofs", src, "sqlite", drv)
	if err != nil {
		return err
	}

	if err := migrator.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
