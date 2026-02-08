package database

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"tech-quest/internal/configs"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

func RunMigrations() error {
	cfg := configs.Configs
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBSSLMode,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("failed to open database connection: %w", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(db)
	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("failed to set goose dialect: %w", err)
	}

	// Определяем путь к миграциям
	migrationsPath := "migrations"
	// Проверяем, есть ли миграции в /migrations (для Docker)
	if _, err := os.Stat("/migrations"); err == nil {
		migrationsPath = "/migrations"
	} else {
		// Проверяем относительный путь от рабочей директории
		if _, err := os.Stat("migrations"); err != nil {
			// Пытаемся найти миграции относительно текущего файла
			_, filename, _, _ := runtime.Caller(0)
			migrationsPath = filepath.Join(filepath.Dir(filename), "../../migrations")
		}
	}

	if err := goose.Up(db, migrationsPath); err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}
	return nil
}
