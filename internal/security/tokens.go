package security

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenTokenForUser(userId string, secret string) string {
	payload := jwt.MapClaims{
		"sub": userId,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		panic(err)
	}

	return t
}

func GetUserFromToken(token string, secret string) (string, error) {
	parsed, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return "", errors.New("invalid token")
	}

	userId, err := parsed.Claims.GetSubject()

	if err != nil || userId == "" {
		return "", errors.New("invalid token")
	}

	return userId, nil
}
