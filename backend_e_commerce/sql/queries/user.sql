-- name: CreateUser :exec
INSERT INTO users (firstName, lastName, email, password)
VALUES (?, ?, ?, ?);

-- name: GetUserByID :one
SELECT id, firstName, lastName, email, password, createdAt
FROM users
WHERE id = ?;

-- name: GetUserByEmail :one
SELECT id, firstName, lastName, email, password, createdAt
FROM users
WHERE email = ?;

-- name: ListUsers :many
SELECT id, firstName, lastName, email, password, createdAt
FROM users
ORDER BY createdAt DESC;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?;
