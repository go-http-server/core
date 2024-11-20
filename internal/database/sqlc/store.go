package database

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Store interface {
	Querier
}

// SQLStore provider all functions to execute SQL queries and transaction
type SQLStore struct {
	*Queries
	connectionPool *pgxpool.Pool
}

func NewStore(db *pgxpool.Pool) Store {
	return &SQLStore{
		Queries:        New(db),
		connectionPool: db,
	}
}
