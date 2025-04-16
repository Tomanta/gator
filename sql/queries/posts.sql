-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES (
    $1,
    NOW(),
    NOW(),
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetPosts :many
SELECT posts.id, posts.created_at, posts.updated_at, posts.title, posts.url, posts.description, posts.published_at
FROM feeds
INNER JOIN feed_follows ff ON ff.feed_id = feeds.id
INNER JOIN posts ON posts.feed_id = feeds.id
WHERE ff.user_id = $1
ORDER BY published_at DESC
LIMIT $2;