package utils

import (
	"chat-application/models"
	"github.com/beego/beego/v2/core/logs"
	"github.com/golang-jwt/jwt/v4"
	_ "github.com/satori/go.uuid"
	"os"
	"time"
)

func GenerateJWTToken(expirationTime time.Time, user *models.Token) (string, []byte) {
	getSecretKey := os.Getenv("SECRET_KEY")
	secretKey := []byte(getSecretKey)
	claims := &models.Claims{
		VerifyID: user.VerifyID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, signErr := token.SignedString(secretKey)
	CheckError(signErr, "Error signing token ")
	return tokenString, secretKey
}

func ValidateJwt(user models.Token, token string, secretKey []byte) error {
	logs.Info("token ", token)
	tokenValid, parseErr := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			logs.Info("Unauthorised user as signing method not matched")
			return nil, nil
		}
		return secretKey, nil
	})
	if parseErr != nil {
		logs.Info("Error parsing token", parseErr)
		return parseErr
	}
	if tokenValid.Valid {
		logs.Info("Valid token: Success")
	} else {
		logs.Info("Invalid token: Unauthorised user")
	}
	return nil
}
