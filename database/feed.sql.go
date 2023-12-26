// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: feed.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const articleFeed = `-- name: ArticleFeed :many
SELECT 
    a.id AS article_id,
    a.title,
    a.link,
    a.created_at,
    u.username,
    COUNT(c.id) AS comment_count
FROM 
    articles a
JOIN 
    users u ON a.user_id = u.id
LEFT JOIN 
    comments c ON a.id = c.article_id
GROUP BY 
    a.id, u.username
ORDER BY 
    a.created_at DESC
LIMIT 
    $1
OFFSET 
    $2
`

type ArticleFeedParams struct {
	Limit  int32
	Offset int32
}

type ArticleFeedRow struct {
	ArticleID    pgtype.UUID
	Title        string
	Link         string
	CreatedAt    pgtype.Timestamptz
	Username     string
	CommentCount int64
}

func (q *Queries) ArticleFeed(ctx context.Context, arg ArticleFeedParams) ([]ArticleFeedRow, error) {
	rows, err := q.db.Query(ctx, articleFeed, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ArticleFeedRow
	for rows.Next() {
		var i ArticleFeedRow
		if err := rows.Scan(
			&i.ArticleID,
			&i.Title,
			&i.Link,
			&i.CreatedAt,
			&i.Username,
			&i.CommentCount,
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

const createArticle = `-- name: CreateArticle :one
INSERT INTO articles (user_id, title, link) VALUES ($1, $2, $3) RETURNING id, user_id, title, link, created_at, updated_at
`

type CreateArticleParams struct {
	UserID pgtype.UUID
	Title  string
	Link   string
}

func (q *Queries) CreateArticle(ctx context.Context, arg CreateArticleParams) (Article, error) {
	row := q.db.QueryRow(ctx, createArticle, arg.UserID, arg.Title, arg.Link)
	var i Article
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Link,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}