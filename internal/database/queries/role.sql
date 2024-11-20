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
