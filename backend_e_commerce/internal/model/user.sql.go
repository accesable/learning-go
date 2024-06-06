// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: user.sql

package mysqlc

import (
	"context"
)

const createUser = `-- name: CreateUser :exec
INSERT INTO users (firstName, lastName, email, password)
VALUES (?, ?, ?, ?)
`

type CreateUserParams struct {
	Firstname string
	Lastname  string
	Email     string
	Password  string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.ExecContext(ctx, createUser,
		arg.Firstname,
		arg.Lastname,
		arg.Email,
		arg.Password,
	)
	return err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?
`

func (q *Queries) DeleteUser(ctx context.Context, id uint32) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, firstName, lastName, email, password, createdAt
FROM users
WHERE email = ?
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Firstname,
		&i.Lastname,
		&i.Email,
		&i.Password,
		&i.Createdat,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, firstName, lastName, email, password, createdAt
FROM users
WHERE id = ?
`

func (q *Queries) GetUserByID(ctx context.Context, id uint32) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Firstname,
		&i.Lastname,
		&i.Email,
		&i.Password,
		&i.Createdat,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, firstName, lastName, email, password, createdAt
FROM users
ORDER BY createdAt DESC
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Firstname,
			&i.Lastname,
			&i.Email,
			&i.Password,
			&i.Createdat,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}