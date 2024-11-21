// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: access_path.sql

package database

import (
	"context"
)

const createAccessPath = `-- name: CreateAccessPath :one
INSERT INTO access_paths (role_id, path_id) VALUES ($1, $2) RETURNING id, role_id, path_id
`

type CreateAccessPathParams struct {
	RoleID int64 `json:"role_id"`
	PathID int64 `json:"path_id"`
}

func (q *Queries) CreateAccessPath(ctx context.Context, arg CreateAccessPathParams) (AccessPath, error) {
	row := q.db.QueryRow(ctx, createAccessPath, arg.RoleID, arg.PathID)
	var i AccessPath
	err := row.Scan(&i.ID, &i.RoleID, &i.PathID)
	return i, err
}

const deleteAccessPath = `-- name: DeleteAccessPath :exec
DELETE FROM access_paths
  WHERE id = $1
`

func (q *Queries) DeleteAccessPath(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteAccessPath, id)
	return err
}

const getAccessPath = `-- name: GetAccessPath :many
SELECT id, role_id, path_id FROM access_paths WHERE role_id = $1
`

func (q *Queries) GetAccessPath(ctx context.Context, roleID int64) ([]AccessPath, error) {
	rows, err := q.db.Query(ctx, getAccessPath, roleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AccessPath{}
	for rows.Next() {
		var i AccessPath
		if err := rows.Scan(&i.ID, &i.RoleID, &i.PathID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}