package User

import (
	"database/sql"
	"time"
)

type User struct {
	PkID       int            `json:"pk_id"`
	UserName   string         `json:"user_name"`
	Password   sql.NullString `json:"password"`
	IsDelete   int            `json:"is_delete"`
	CreateTime time.Time      `json:"create_time"`
	UpdateTime sql.NullTime   `json:"update_time"`
	Phone      sql.NullString `json:"phone"`
	Email      sql.NullString `json:"email"`
}
