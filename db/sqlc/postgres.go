package db

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitDB(dsn string) *pgxpool.Pool {
	log.Println("Starting DB...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatal("failed to parse DATABASE_URL:", err)
	}

	// ✅ Pool tuning (safe defaults)
	config.MaxConns = 10
	config.MinConns = 2
	config.MaxConnIdleTime = 5 * time.Minute
	config.MaxConnLifetime = 30 * time.Minute

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Fatal("unable to create pgx pool:", err)
	}

	// Do NOT crash app if DB is temporarily unavailable
	if err := pool.Ping(ctx); err != nil {
		log.Println("DB not reachable at startup:", err)
	}

	log.Println("DB pool initialized")

	return pool
}
