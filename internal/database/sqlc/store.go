package database

import (
	"context"
	"fmt"

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

func (sqlStore *SQLStore) execTX(ctx context.Context, fn func(*Queries) error) error {
	transaction, err := sqlStore.connectionPool.Begin(ctx)
	if err != nil {
		return err
	}

	queries := New(transaction)
	err = fn(queries)
	if err != nil {
		if errRollback := transaction.Rollback(ctx); errRollback != nil {
			return fmt.Errorf("[ERROR - Rollback]: %s", errRollback)
		}
		return fmt.Errorf("[ERROR - Unknow]: %s", err)
	}

	errCommit := transaction.Commit(ctx)
	if errCommit != nil {
		return fmt.Errorf("[ERROR - Commit]: %s", errCommit)
	}

	return nil
}
