// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"context"
)

type Querier interface {
	CreatePath(ctx context.Context, arg CreatePathParams) (Path, error)
	ListPaths(ctx context.Context, arg ListPathsParams) ([]Path, error)
}

var _ Querier = (*Queries)(nil)
