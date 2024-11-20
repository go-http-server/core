-- name: ListPaths :many
SELECT * FROM paths
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: CreatePath :one
INSERT INTO paths (
  path_name, path, path_description
) values (
  $1, $2, $3
) RETURNING *;
