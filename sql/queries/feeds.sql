-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeeds :many
SELECT feeds.name, feeds.url, users.name AS creator
FROM feeds
INNER JOIN users ON feeds.user_id = users.id;

-- name: GetFeedByURL :one
SELECt feeds.name, feeds.url, feeds.id
FROM feeds
WHERE feeds.url = $1;

-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES (
        $1, $2, $3, $4, $5
    )
    RETURNING *
)
SELECT
    inserted_feed_follow.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feed_follow
INNER JOIN feeds ON inserted_feed_follow.feed_id = feeds.id
INNER JOIN users ON inserted_feed_follow.user_id = users.id;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows
WHERE feed_id = $1
  AND user_id = $2;

-- name: GetFeedFollowsForUser :many
SELECT feeds.name AS feed_name, feeds.url AS feed_url, users.name AS user_name
FROM feed_follows
INNER JOIN feeds ON feed_follows.feed_id = feeds.id
INNER JOIN users ON feed_follows.user_id = users.id
WHERE feed_follows.user_id = $1;

-- name: MarkFeedFetched :exec
UPDATE feeds
SET last_fetched_at = $1, updated_at = $1
WHERE feeds.id = $2;

-- name: GetNextFeedToFetch :one
SELECT feeds.name, feeds.url, feeds.id, feeds.last_fetched_at
FROM feeds
ORDER BY last_fetched_at DESC NULLS FIRST
LIMIT 1;