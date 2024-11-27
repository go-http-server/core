-- name: GetUser :one
SELECT * FROM users WHERE username = $1 OR email = $2;

-- name: CreateUser :one
INSERT INTO users (
  username, hashed_password, email, full_name, role_id, code_verify_email
) VALUES ($1, $2, $3, $4, $5, $6)
  RETURNING *;

-- name: UpdateUser :one
UPDATE users
  SET code_verify_email = coalesce(sqlc.narg(code_verify_email), code_verify_email),
      code_reset_password = coalesce(sqlc.narg(code_reset_password), code_reset_password),
      hashed_password = coalesce(sqlc.narg(hashed_password), hashed_password),
      email = coalesce(sqlc.narg(email), email),
      is_verified_email= coalesce(sqlc.narg(is_verified_email), is_verified_email),
      full_name = coalesce(sqlc.narg(full_name), full_name),
      role_id = coalesce(sqlc.narg(role_id), role_id),
      token = coalesce(sqlc.narg(token), token)
  WHERE username = $1 RETURNING *;
