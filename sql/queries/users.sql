-- name: CreateUsers :execresult
INSERT INTO users (create_at, update_at, name) VALUES (?, ?, ?);

-- name: LastInsertedUserID :one
SELECT LAST_INSERT_ID() AS id;

-- name: GetUserInfo :one
SELECT * from users where id = ?;