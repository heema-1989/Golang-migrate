// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package users

import (
	"database/sql"
	"time"

	"github.com/satori/go.uuid"
)

type Room struct {
	ID       int64  `json:"id"`
	RoomName string `json:"room_name"`
	UserID   int64  `json:"user_id"`
}

type User struct {
	ID                 int64        `json:"id"`
	VerifyID           uuid.UUID    `json:"verify_id"`
	FullName           string       `json:"full_name"`
	UserName           string       `json:"user_name"`
	Email              string       `json:"email"`
	Password           string       `json:"password"`
	CreatedAt          time.Time    `json:"created_at"`
	PasswordCreateDate time.Time    `json:"password_create_date"`
	PasswordUpdateDate sql.NullTime `json:"password_update_date"`
}
