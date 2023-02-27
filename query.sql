-- name: GetBlogById :one
SELECT * FROM blogs
WHERE id = $1 LIMIT 1;

-- name: GetBlogs :many
SELECT * FROM blogs
ORDER BY create_date DESC;

-- name: CreateBlog :one
INSERT INTO blogs (
  title, content, create_date
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: DeleteBlog :exec
DELETE FROM blogs
WHERE id = $1;

-- name: UpdateBlog :one
UPDATE blogs
  set title = $2, content = $3
WHERE id = $1
RETURNING *;

-- name: GetChatById :one
SELECT * FROM chat_history
WHERE id = $1 LIMIT 1;

-- name: GetChatHistory :many
SELECT * FROM chat_history
ORDER BY create_date DESC;

-- name: CreateChat :one
INSERT INTO chat_history (
  sender, message, create_date
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: DeleteChat :exec
DELETE FROM chat_history
WHERE id = $1;