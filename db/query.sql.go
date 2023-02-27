// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: query.sql

package queries

import (
	"context"
	"time"
)

const createBlog = `-- name: CreateBlog :one
INSERT INTO blogs (
  title, content, create_date
) VALUES (
  $1, $2, $3
) RETURNING id, title, content, create_date
`

type CreateBlogParams struct {
	Title      string
	Content    string
	CreateDate time.Time
}

func (q *Queries) CreateBlog(ctx context.Context, arg CreateBlogParams) (Blog, error) {
	row := q.db.QueryRowContext(ctx, createBlog, arg.Title, arg.Content, arg.CreateDate)
	var i Blog
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.CreateDate,
	)
	return i, err
}

const createChat = `-- name: CreateChat :one
INSERT INTO chat_history (
  sender, message, create_date
) VALUES (
  $1, $2, $3
) RETURNING id, sender, message, create_date
`

type CreateChatParams struct {
	Sender     string
	Message    string
	CreateDate time.Time
}

func (q *Queries) CreateChat(ctx context.Context, arg CreateChatParams) (ChatHistory, error) {
	row := q.db.QueryRowContext(ctx, createChat, arg.Sender, arg.Message, arg.CreateDate)
	var i ChatHistory
	err := row.Scan(
		&i.ID,
		&i.Sender,
		&i.Message,
		&i.CreateDate,
	)
	return i, err
}

const deleteBlog = `-- name: DeleteBlog :exec
DELETE FROM blogs
WHERE id = $1
`

func (q *Queries) DeleteBlog(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteBlog, id)
	return err
}

const deleteChat = `-- name: DeleteChat :exec
DELETE FROM chat_history
WHERE id = $1
`

func (q *Queries) DeleteChat(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteChat, id)
	return err
}

const getBlogById = `-- name: GetBlogById :one
SELECT id, title, content, create_date FROM blogs
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetBlogById(ctx context.Context, id int64) (Blog, error) {
	row := q.db.QueryRowContext(ctx, getBlogById, id)
	var i Blog
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.CreateDate,
	)
	return i, err
}

const getBlogs = `-- name: GetBlogs :many
SELECT id, title, content, create_date FROM blogs
ORDER BY create_date DESC
`

func (q *Queries) GetBlogs(ctx context.Context) ([]Blog, error) {
	rows, err := q.db.QueryContext(ctx, getBlogs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Blog
	for rows.Next() {
		var i Blog
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Content,
			&i.CreateDate,
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

const getChatById = `-- name: GetChatById :one
SELECT id, sender, message, create_date FROM chat_history
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetChatById(ctx context.Context, id int64) (ChatHistory, error) {
	row := q.db.QueryRowContext(ctx, getChatById, id)
	var i ChatHistory
	err := row.Scan(
		&i.ID,
		&i.Sender,
		&i.Message,
		&i.CreateDate,
	)
	return i, err
}

const getChatHistory = `-- name: GetChatHistory :many
SELECT id, sender, message, create_date FROM chat_history
ORDER BY create_date DESC
`

func (q *Queries) GetChatHistory(ctx context.Context) ([]ChatHistory, error) {
	rows, err := q.db.QueryContext(ctx, getChatHistory)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ChatHistory
	for rows.Next() {
		var i ChatHistory
		if err := rows.Scan(
			&i.ID,
			&i.Sender,
			&i.Message,
			&i.CreateDate,
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

const updateBlog = `-- name: UpdateBlog :one
UPDATE blogs
  set title = $2, content = $3
WHERE id = $1
RETURNING id, title, content, create_date
`

type UpdateBlogParams struct {
	ID      int64
	Title   string
	Content string
}

func (q *Queries) UpdateBlog(ctx context.Context, arg UpdateBlogParams) (Blog, error) {
	row := q.db.QueryRowContext(ctx, updateBlog, arg.ID, arg.Title, arg.Content)
	var i Blog
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.CreateDate,
	)
	return i, err
}