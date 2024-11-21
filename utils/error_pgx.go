package utils

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

const (
	ForeignKeyViolation = "23503"
	UniqueViolation     = "23505"
)

func ErrorCodePgxConstraint(err error) string {
	var pgxError *pgconn.PgError
	if errors.As(err, &pgxError) {
		return pgxError.Code
	}
	return ""
}
