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

-- name: GetOnePath :one
SELECT * FROM paths WHERE id = $1;

-- name: UpdatePath :one
UPDATE paths
  SET path_name = coalesce(sqlc.narg(path_name), path_name),
      path = coalesce(sqlc.narg(path), path),
      path_description = coalesce(sqlc.narg(path_description), path_description)
  WHERE id = $1
  RETURNING *;
