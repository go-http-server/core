// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"context"
)

type Querier interface {
	CreateAccessPath(ctx context.Context, arg CreateAccessPathParams) (AccessPath, error)
	CreatePath(ctx context.Context, arg CreatePathParams) (Path, error)
	CreateRole(ctx context.Context, arg CreateRoleParams) (Role, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteAccessPath(ctx context.Context, id int64) error
	GetAccessPath(ctx context.Context, roleID int64) ([]AccessPath, error)
	GetOnePath(ctx context.Context, id int64) (Path, error)
	GetOneRole(ctx context.Context, id int64) (Role, error)
	GetRoles(ctx context.Context, arg GetRolesParams) ([]Role, error)
	GetUser(ctx context.Context, arg GetUserParams) (User, error)
	ListPaths(ctx context.Context, arg ListPathsParams) ([]Path, error)
	UpdatePath(ctx context.Context, arg UpdatePathParams) (Path, error)
	UpdateRole(ctx context.Context, arg UpdateRoleParams) (Role, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)
