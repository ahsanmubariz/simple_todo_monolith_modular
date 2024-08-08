package migrations

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"

	"github.com/ahsanmubariz/simple_todo_monolith_modular/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations() error {
	migrationDir, err := getMigrationDirectory()
	if err != nil {
		fmt.Println("Error getting migration directory:", err)
		return err
	}

	migrationURL := fmt.Sprintf("file://%s", migrationDir)
	fmt.Println("Migration URL:", migrationURL)

	// Check if the directory and files exist and are accessible
	if err := checkMigrationFiles(migrationDir); err != nil {
		fmt.Println("Error checking migration files:", err)
		return err
	}

	databaseDSN := convertToURLDSN(config.GetDatabaseDSN())
	fmt.Println("Converted Database DSN:", databaseDSN)

	m, err := migrate.New(
		migrationURL,
		databaseDSN,
	)
	if err != nil {
		fmt.Println("Migrate New error:", err)
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		fmt.Println("Migrate Up error:", err)
		return err
	}

	return nil
}

func getMigrationDirectory() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	fmt.Println("Current working directory:", dir)

	migrationPath := filepath.Join(dir, "migrations")
	fmt.Println("Constructed migration path:", migrationPath)

	if _, err := os.Stat(migrationPath); os.IsNotExist(err) {
		fmt.Println("Migration directory does not exist:", migrationPath)
		return "", err
	}

	return migrationPath, nil
}

func checkMigrationFiles(migrationDir string) error {
	files, err := os.ReadDir(migrationDir)
	if err != nil {
		return err
	}

	fmt.Println("Files in migration directory:")
	for _, file := range files {
		fmt.Println(" -", file.Name())
	}

	return nil
}

func convertToURLDSN(dsn string) string {
	var host, user, password, dbname, port, sslmode string
	fmt.Sscanf(dsn, "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", &host, &user, &password, &dbname, &port, &sslmode)

	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		url.QueryEscape(user),
		url.QueryEscape(password),
		host,
		port,
		dbname,
		sslmode,
	)
}
