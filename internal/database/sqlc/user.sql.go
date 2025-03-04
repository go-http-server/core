// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  username, hashed_password, email, full_name, role_id, code_verify_email
) VALUES ($1, $2, $3, $4, $5, $6)
  RETURNING username, code_verify_email, code_reset_password, created_at, hashed_password, email, is_verified_email, full_name, role_id, token, password_changed_at
`

type CreateUserParams struct {
	Username        string      `json:"username"`
	HashedPassword  string      `json:"hashed_password"`
	Email           string      `json:"email"`
	FullName        string      `json:"full_name"`
	RoleID          int64       `json:"role_id"`
	CodeVerifyEmail pgtype.Text `json:"code_verify_email"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Username,
		arg.HashedPassword,
		arg.Email,
		arg.FullName,
		arg.RoleID,
		arg.CodeVerifyEmail,
	)
	var i User
	err := row.Scan(
		&i.Username,
		&i.CodeVerifyEmail,
		&i.CodeResetPassword,
		&i.CreatedAt,
		&i.HashedPassword,
		&i.Email,
		&i.IsVerifiedEmail,
		&i.FullName,
		&i.RoleID,
		&i.Token,
		&i.PasswordChangedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT username, code_verify_email, code_reset_password, created_at, hashed_password, email, is_verified_email, full_name, role_id, token, password_changed_at FROM users WHERE username = $1 OR email = $1
`

func (q *Queries) GetUser(ctx context.Context, identifier string) (User, error) {
	row := q.db.QueryRow(ctx, getUser, identifier)
	var i User
	err := row.Scan(
		&i.Username,
		&i.CodeVerifyEmail,
		&i.CodeResetPassword,
		&i.CreatedAt,
		&i.HashedPassword,
		&i.Email,
		&i.IsVerifiedEmail,
		&i.FullName,
		&i.RoleID,
		&i.Token,
		&i.PasswordChangedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
  SET code_verify_email = coalesce($2, code_verify_email),
      code_reset_password = coalesce($3, code_reset_password),
      hashed_password = coalesce($4, hashed_password),
      email = coalesce($5, email),
      is_verified_email= coalesce($6, is_verified_email),
      full_name = coalesce($7, full_name),
      role_id = coalesce($8, role_id),
      token = coalesce($9, token)
  WHERE username = $1 RETURNING username, code_verify_email, code_reset_password, created_at, hashed_password, email, is_verified_email, full_name, role_id, token, password_changed_at
`

type UpdateUserParams struct {
	Username          string      `json:"username"`
	CodeVerifyEmail   pgtype.Text `json:"code_verify_email"`
	CodeResetPassword pgtype.Text `json:"code_reset_password"`
	HashedPassword    pgtype.Text `json:"hashed_password"`
	Email             pgtype.Text `json:"email"`
	IsVerifiedEmail   pgtype.Bool `json:"is_verified_email"`
	FullName          pgtype.Text `json:"full_name"`
	RoleID            pgtype.Int8 `json:"role_id"`
	Token             pgtype.Text `json:"token"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUser,
		arg.Username,
		arg.CodeVerifyEmail,
		arg.CodeResetPassword,
		arg.HashedPassword,
		arg.Email,
		arg.IsVerifiedEmail,
		arg.FullName,
		arg.RoleID,
		arg.Token,
	)
	var i User
	err := row.Scan(
		&i.Username,
		&i.CodeVerifyEmail,
		&i.CodeResetPassword,
		&i.CreatedAt,
		&i.HashedPassword,
		&i.Email,
		&i.IsVerifiedEmail,
		&i.FullName,
		&i.RoleID,
		&i.Token,
		&i.PasswordChangedAt,
	)
	return i, err
}
