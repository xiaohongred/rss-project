// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0

package database

import (
	"time"
)

type User struct {
	// 主键
	ID int32
	// 创建时间
	CreateAt time.Time
	// 修改时间
	UpdateAt time.Time
	// 姓名
	Name string
}
