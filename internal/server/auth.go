package server

import (
	"errors"
	"strings"

	"github.com/garnet-geo/garnet-userapi/internal/consts"
	"github.com/garnet-geo/garnet-userapi/internal/env"
	"github.com/garnet-geo/garnet-userapi/internal/security"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := extractAuthToken(ctx)
		if err != nil {
			ctx.JSON(401,
				"unauthorized",
			)
			ctx.Abort()
			return
		}

		userID, err := security.GetUserFromToken(token, env.GetSecurityTokenSecret())
		if err != nil {
			ctx.JSON(401,
				"unauthorized",
			)
			ctx.Abort()
			return
		}

		ctx.Set(consts.UserIDContextKey, userID)
	}
}

func extractAuthToken(ctx *gin.Context) (string, error) {
	header := ctx.GetHeader(consts.AuthorizationHeader)
	if header == "" {
		return "", errors.New("request has no auth header")
	}

	parts := strings.Split(header, " ")
	if len(parts) != 2 {
		return "", errors.New("request has no auth header")
	}

	if parts[0] != consts.BearerTokenPrefix {
		return "", errors.New("request has no auth header")
	}

	return parts[1], nil
}
