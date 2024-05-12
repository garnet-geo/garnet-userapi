package security

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenTokenForUser(userId string, secret string) string {
	payload := jwt.MapClaims{
		"user": userId,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		panic(err)
	}

	return t
}
