package database

import (
	"fmt"
	"tech-quest/internal/configs"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewDB() (*sqlx.DB, error) {
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
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}
	return db, nil
}
