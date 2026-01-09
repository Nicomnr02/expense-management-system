package database

import (
	"context"
	"expense-management-system/config"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

type Database struct {
	Cfg  *config.Config
	Conn *pgxpool.Pool
}

func New(Cfg *config.Config) *Database {
	return &Database{Cfg, nil}
}

func (d *Database) Init() error {
	if os.Getenv("APP_ENV") == "" {
		err := godotenv.Load()
		if err != nil {
			return err
			// log.Fatal("No .env loaded")
		}
	}

	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		d.Cfg.DBUser, d.Cfg.DBPassword, d.Cfg.DBHost, d.Cfg.DBPort, d.Cfg.DBName,
	)

	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return err
		// log.Fatalf("Unable to parse config: %v", err)
	}

	config.MaxConns = int32(d.Cfg.DBMaxConns)
	config.MaxConnIdleTime = time.Duration(d.Cfg.DBMaxConnIdle) * time.Minute

	ctx := context.Background()

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return err
		// log.Fatalf("Unable to connect to database: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, time.Duration(d.Cfg.DBConnectionTimeout)*time.Second)
	defer cancel()
	if err := pool.Ping(ctx); err != nil {
		return err
		// log.Fatalf("Unable to ping DB: %v", err)
	}

	d.Conn = pool

	return nil
}

func (d *Database) Close() {
	d.Conn.Close()
}
