package db

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"log"
	"os"
	"path/filepath"
)

func readFile(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

type migration struct {
	Name         string
	upFilePath   string
	downFilePath string
}

func (m *migration) GetUp() (string, error) {
	return readFile(m.upFilePath)
}

func (m *migration) GetDown() (string, error) {
	return readFile(m.downFilePath)
}

func readMigrations(migrationPath string) (map[string]migration, error) {
	migrations := make(map[string]migration)

	files, err := os.ReadDir(migrationPath)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			migrationName := file.Name()
			upFilePath := filepath.Join(migrationPath, migrationName, "up.sql")
			downFilePath := filepath.Join(migrationPath, migrationName, "down.sql")

			migrations[migrationName] = migration{migrationName, upFilePath, downFilePath}
		}
	}

	return migrations, nil
}

func migrateUp(tx pgx.Tx, migration *migration, ctx *context.Context) error {
	upSQL, err := migration.GetUp()
	if err != nil {
		return err
	}

	_, err = tx.Exec(*ctx, upSQL)
	if err != nil {
		return err
	}

	log.Printf("Applied migration: %s\n", migration.Name)
	return nil
}

const noTableErrorCode = "42P01"

func createMigrationTableIfNotExists(conn *pgx.Conn, ctx *context.Context) error {
	_, err := conn.Exec(*ctx, "SELECT 1 FROM _migrations LIMIT 1")
	if err == nil {
		return nil
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == noTableErrorCode {
		tableSql := "CREATE TABLE _migrations (id SERIAL PRIMARY KEY, name TEXT NOT NULL)"
		_, err = conn.Exec(*ctx, tableSql)
		if err != nil {
			return err
		}
		return nil
	} else {
		return err
	}
}

func runMigrations(conn *pgx.Conn, migrations map[string]migration, ctx *context.Context) error {
	err := createMigrationTableIfNotExists(conn, ctx)
	if err != nil {
		return err
	}

	rows, err := conn.Query(*ctx, "SELECT name FROM _migrations ORDER BY name")
	if err != nil {
		return err
	}
	defer rows.Close()

	appliedMigrations := make(map[string]bool)
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			return err
		}
		appliedMigrations[name] = true
	}

	if len(appliedMigrations) == len(migrations) {
		log.Printf("No migrations to run")
		return nil
	}

	return pgx.BeginFunc(*ctx, conn, func(tx pgx.Tx) error {
		for _, migration := range migrations {
			if !appliedMigrations[migration.Name] {
				err := migrateUp(tx, &migration, ctx)
				if err != nil {
					return err
				}

				insertSQL := "INSERT INTO _migrations (name) VALUES ($1)"
				_, err = tx.Exec(*ctx, insertSQL, migration.Name)
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
}

func RunMigrations(migrationsPath string) error {
	log.Printf("Running migrations...")
	ctx := context.Background()
	migrations, err := readMigrations(migrationsPath)
	if err != nil {
		return err
	}

	err = runMigrations(Conn, migrations, &ctx)
	if err != nil {
		log.Printf("Rollback all previous transactions")
		return err
	}

	return nil
}
