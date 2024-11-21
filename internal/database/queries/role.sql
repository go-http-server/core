-- name: GetRoles :many
SELECT * FROM roles
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: CreateRole :one
INSERT INTO roles (
  role_name, description
) VALUES ($1, $2)
RETURNING *;

-- name: UpdateRole :one
UPDATE roles
SET role_name = coalesce(sqlc.narg(role_name), role_name),
    description = coalesce(sqlc.narg(description), description)
WHERE id = $1
RETURNING *;
