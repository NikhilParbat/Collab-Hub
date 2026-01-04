package db

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

// Store is an interface for executing queries and transactions related to rooms and peers.
type Store interface {
	Querier
}

// SQLStore provides methods to execute SQL queries.
type SQLStore struct {
	*Queries
	db *pgxpool.Pool
}

// NewStore creates a new SQLStore instance with the given database connection pool.
func NewStore(db *pgxpool.Pool) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}
