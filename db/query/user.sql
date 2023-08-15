-- name: CreateUser :one
INSERT INTO users (
  username ,fullname, email, password
) VALUES (
  $1, $2, $3, $4
) RETURNING id, username ,fullname, email, created_at;

-- name: GetUser :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: GetUserById :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;