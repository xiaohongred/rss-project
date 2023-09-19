-- name: CreatePost :one
insert into posts (id, created_at, updated_at, title, description, published_at, url, feed_id)
values ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;
--

-- name: GetPostsForUser :many
SELECT posts.* FROM posts
JOIN feed_follows ON feed_follows.feed_id = posts.feed_id
WHERE feed_follows.user_id = $1
ORDER BY posts.published_at DESC
LIMIT $2;
--