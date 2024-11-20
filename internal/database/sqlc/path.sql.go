// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: path.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createPath = `-- name: CreatePath :one
INSERT INTO paths (
  path_name, path, path_description
) values (
  $1, $2, $3
) RETURNING id, path_name, path, path_description
`

type CreatePathParams struct {
	PathName        string      `json:"path_name"`
	Path            string      `json:"path"`
	PathDescription pgtype.Text `json:"path_description"`
}

func (q *Queries) CreatePath(ctx context.Context, arg CreatePathParams) (Path, error) {
	row := q.db.QueryRow(ctx, createPath, arg.PathName, arg.Path, arg.PathDescription)
	var i Path
	err := row.Scan(
		&i.ID,
		&i.PathName,
		&i.Path,
		&i.PathDescription,
	)
	return i, err
}

const listPaths = `-- name: ListPaths :many
SELECT id, path_name, path, path_description FROM paths
ORDER BY id
`

func (q *Queries) ListPaths(ctx context.Context) ([]Path, error) {
	rows, err := q.db.Query(ctx, listPaths)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Path{}
	for rows.Next() {
		var i Path
		if err := rows.Scan(
			&i.ID,
			&i.PathName,
			&i.Path,
			&i.PathDescription,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
