// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package database

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Article struct {
	ID        pgtype.UUID
	UserID    pgtype.UUID
	Title     string
	Link      string
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
}

type Comment struct {
	ID        pgtype.UUID
	ArticleID pgtype.UUID
	UserID    pgtype.UUID
	Content   string
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
}

type Like struct {
	ID        pgtype.UUID
	ArticleID pgtype.UUID
	UserID    pgtype.UUID
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
}

type User struct {
	ID          pgtype.UUID
	Username    string
	Email       string
	Password    string
	ConfirmedAt pgtype.Timestamptz
	CreatedAt   pgtype.Timestamptz
	UpdatedAt   pgtype.Timestamptz
}

type UserToken struct {
	ID        pgtype.UUID
	UserID    pgtype.UUID
	Token     string
	Context   string
	CreatedAt pgtype.Timestamptz
}
