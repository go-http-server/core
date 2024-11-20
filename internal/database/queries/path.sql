-- name: ListPaths :many
SELECT * FROM paths
ORDER BY id;

-- name: CreatePath :one
INSERT INTO paths (
  path_name, path, path_description
) values (
  $1, $2, $3
) RETURNING *;
