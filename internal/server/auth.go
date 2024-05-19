package server

import (
	"errors"
	"strings"

	"github.com/garnet-geo/garnet-userapi/internal/consts"
	"github.com/garnet-geo/garnet-userapi/internal/env"
	"github.com/garnet-geo/garnet-userapi/internal/security"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Debugln("Handling authorization for", ctx.Request.URL.Path)

		token, err := extractAuthToken(ctx)
		if err != nil {
			log.Debugln("Err extracting auth token", err)
			ctx.JSON(401, "unauthorized")
			ctx.Abort()
			return
		}

		userID, err := security.GetUserFromToken(token, env.GetSecurityTokenSecret())
		if err != nil {
			log.Debugln("Err getting user from token", err)
			ctx.JSON(401, "unauthorized")
			ctx.Abort()
			return
		}

		log.Debugln("Got user from token", userID)
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
		return "", errors.New("2 parts of auth header expected")
	}

	if parts[0] != consts.BearerTokenPrefix {
		return "", errors.New("first part have to be Bearer")
	}

	return parts[1], nil
}
