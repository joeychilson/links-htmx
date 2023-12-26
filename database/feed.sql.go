// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: feed.sql

package database

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const countVotes = `-- name: CountVotes :one
SELECT COUNT(*) FROM votes WHERE link_id = $1
`

func (q *Queries) CountVotes(ctx context.Context, linkID uuid.UUID) (int64, error) {
	row := q.db.QueryRow(ctx, countVotes, linkID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createLink = `-- name: CreateLink :exec
INSERT INTO links (user_id, title, url) VALUES ($1, $2, $3)
`

type CreateLinkParams struct {
	UserID uuid.UUID
	Title  string
	Url    string
}

func (q *Queries) CreateLink(ctx context.Context, arg CreateLinkParams) error {
	_, err := q.db.Exec(ctx, createLink, arg.UserID, arg.Title, arg.Url)
	return err
}

const createVote = `-- name: CreateVote :exec
INSERT INTO votes (user_id, link_id) VALUES ($1, $2)
`

type CreateVoteParams struct {
	UserID uuid.UUID
	LinkID uuid.UUID
}

func (q *Queries) CreateVote(ctx context.Context, arg CreateVoteParams) error {
	_, err := q.db.Exec(ctx, createVote, arg.UserID, arg.LinkID)
	return err
}

const linkFeed = `-- name: LinkFeed :many
SELECT 
    l.id AS id,
    l.title,
    l.url,
    l.created_at,
    u.username,
    COUNT(DISTINCT c.id) AS comment_count,
    COUNT(DISTINCT l.id) AS vote_count
FROM 
    links l
JOIN 
    users u ON l.user_id = u.id
LEFT JOIN 
    comments c ON l.id = c.link_id
LEFT JOIN 
    votes v ON l.id = v.link_id
GROUP BY 
    l.id, u.username
ORDER BY 
    l.created_at DESC
LIMIT 
    $1
OFFSET 
    $2
`

type LinkFeedParams struct {
	Limit  int32
	Offset int32
}

type LinkFeedRow struct {
	ID           uuid.UUID
	Title        string
	Url          string
	CreatedAt    pgtype.Timestamptz
	Username     string
	CommentCount int64
	VoteCount    int64
}

func (q *Queries) LinkFeed(ctx context.Context, arg LinkFeedParams) ([]LinkFeedRow, error) {
	rows, err := q.db.Query(ctx, linkFeed, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []LinkFeedRow
	for rows.Next() {
		var i LinkFeedRow
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Url,
			&i.CreatedAt,
			&i.Username,
			&i.CommentCount,
			&i.VoteCount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const userVoted = `-- name: UserVoted :one
SELECT EXISTS(SELECT 1 FROM votes WHERE user_id = $1 AND link_id = $2)
`

type UserVotedParams struct {
	UserID uuid.UUID
	LinkID uuid.UUID
}

func (q *Queries) UserVoted(ctx context.Context, arg UserVotedParams) (bool, error) {
	row := q.db.QueryRow(ctx, userVoted, arg.UserID, arg.LinkID)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}
