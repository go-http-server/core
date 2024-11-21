-- name: GetAccessPath :many
SELECT * FROM access_paths WHERE role_id = $1;

-- name: CreateAccessPath :one
INSERT INTO access_paths (role_id, path_id) VALUES ($1, $2) RETURNING *;

-- name: DeleteAccessPath :exec
DELETE FROM access_paths
  WHERE id = $1;
