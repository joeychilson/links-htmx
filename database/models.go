package database

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Comment struct {
	ID        uuid.UUID
	LinkID    uuid.UUID
	UserID    uuid.UUID
	Content   string
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
}

type Like struct {
	ID        uuid.UUID
	LinkID    uuid.UUID
	UserID    uuid.UUID
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
}

type Link struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Title     string
	Url       string
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
}

type User struct {
	ID          uuid.UUID
	Username    string
	Email       string
	Password    string
	ConfirmedAt pgtype.Timestamptz
	CreatedAt   pgtype.Timestamptz
	UpdatedAt   pgtype.Timestamptz
}

type UserToken struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Token     string
	Context   string
	CreatedAt pgtype.Timestamptz
}
