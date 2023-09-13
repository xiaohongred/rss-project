// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: users.sql

package database

import (
	"context"
	"database/sql"
	"time"
)

const createUsers = `-- name: CreateUsers :execresult
INSERT INTO users (create_at, update_at, name) VALUES (?, ?, ?)
`

type CreateUsersParams struct {
	CreateAt time.Time
	UpdateAt time.Time
	Name     string
}

func (q *Queries) CreateUsers(ctx context.Context, arg CreateUsersParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createUsers, arg.CreateAt, arg.UpdateAt, arg.Name)
}

const getUserInfo = `-- name: GetUserInfo :one
SELECT id, create_at, update_at, name from users where id = ?
`

func (q *Queries) GetUserInfo(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserInfo, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreateAt,
		&i.UpdateAt,
		&i.Name,
	)
	return i, err
}

const lastInsertedUserID = `-- name: LastInsertedUserID :one
SELECT LAST_INSERT_ID() AS id
`

func (q *Queries) LastInsertedUserID(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, lastInsertedUserID)
	var id int64
	err := row.Scan(&id)
	return id, err
}
