// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type AccessPath struct {
	ID     int64 `json:"id"`
	RoleID int64 `json:"role_id"`
	PathID int64 `json:"path_id"`
}

type Path struct {
	ID              int64       `json:"id"`
	PathName        string      `json:"path_name"`
	Path            string      `json:"path"`
	PathDescription pgtype.Text `json:"path_description"`
}

type Role struct {
	ID          int64       `json:"id"`
	RoleName    string      `json:"role_name"`
	Description pgtype.Text `json:"description"`
}

type User struct {
	Username          string           `json:"username"`
	CodeVerifyEmail   pgtype.Text      `json:"code_verify_email"`
	CodeResetPassword pgtype.Text      `json:"code_reset_password"`
	CreatedAt         pgtype.Timestamp `json:"created_at"`
	HashedPassword    string           `json:"hashed_password"`
	Email             string           `json:"email"`
	IsVerifiedEmail   bool             `json:"is_verified_email"`
	FullName          string           `json:"full_name"`
	RoleID            int64            `json:"role_id"`
	Token             pgtype.Text      `json:"token"`
	PasswordChangedAt pgtype.Timestamp `json:"password_changed_at"`
}
