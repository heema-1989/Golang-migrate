package models

import (
	users "chat-application/sqlc-models"
	"github.com/golang-jwt/jwt/v4"
	"github.com/satori/go.uuid"
)

type Claims struct {
	VerifyID uuid.UUID `json:"verify-id"`
	jwt.RegisteredClaims
}
type Token struct {
	JwtToken string
	users.User
}
