// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: users.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const checkEmailExists = `-- name: CheckEmailExists :one
SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)
`

func (q *Queries) CheckEmailExists(ctx context.Context, email string) (bool, error) {
	row := q.db.QueryRow(ctx, checkEmailExists, email)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const checkUsernameExists = `-- name: CheckUsernameExists :one
SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)
`

func (q *Queries) CheckUsernameExists(ctx context.Context, username string) (bool, error) {
	row := q.db.QueryRow(ctx, checkUsernameExists, username)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id
`

type CreateUserParams struct {
	Username string
	Email    string
	Password string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (int32, error) {
	row := q.db.QueryRow(ctx, createUser, arg.Username, arg.Email, arg.Password)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const createUserToken = `-- name: CreateUserToken :one
INSERT INTO user_tokens (user_id, token, context) VALUES ($1, $2, $3) RETURNING token
`

type CreateUserTokenParams struct {
	UserID  int32
	Token   string
	Context string
}

func (q *Queries) CreateUserToken(ctx context.Context, arg CreateUserTokenParams) (string, error) {
	row := q.db.QueryRow(ctx, createUserToken, arg.UserID, arg.Token, arg.Context)
	var token string
	err := row.Scan(&token)
	return token, err
}

const deleteUserToken = `-- name: DeleteUserToken :exec
DELETE FROM user_tokens WHERE token = $1 AND context = $2
`

type DeleteUserTokenParams struct {
	Token   string
	Context string
}

func (q *Queries) DeleteUserToken(ctx context.Context, arg DeleteUserTokenParams) error {
	_, err := q.db.Exec(ctx, deleteUserToken, arg.Token, arg.Context)
	return err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, username, email, password, confirmed_at, created_at, updated_at FROM users WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.ConfirmedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, username, email, confirmed_at, created_at, updated_at FROM users WHERE id = $1
`

type GetUserByIDRow struct {
	ID          int32
	Username    string
	Email       string
	ConfirmedAt pgtype.Timestamptz
	CreatedAt   pgtype.Timestamptz
	UpdatedAt   pgtype.Timestamptz
}

func (q *Queries) GetUserByID(ctx context.Context, id int32) (GetUserByIDRow, error) {
	row := q.db.QueryRow(ctx, getUserByID, id)
	var i GetUserByIDRow
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.ConfirmedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserIDFromToken = `-- name: GetUserIDFromToken :one
SELECT user_id FROM user_tokens WHERE token = $1 AND context = $2
`

type GetUserIDFromTokenParams struct {
	Token   string
	Context string
}

func (q *Queries) GetUserIDFromToken(ctx context.Context, arg GetUserIDFromTokenParams) (int32, error) {
	row := q.db.QueryRow(ctx, getUserIDFromToken, arg.Token, arg.Context)
	var user_id int32
	err := row.Scan(&user_id)
	return user_id, err
}
