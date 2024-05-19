package server

import (
	"fmt"

	"github.com/garnet-geo/garnet-userapi/internal/env"
	"github.com/garnet-geo/garnet-userapi/internal/handlers"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func InitServer() {
	router := gin.Default()

	// Creating group that checks auth token
	authGroup := router.Group("")
	authGroup.Use(AuthMiddleware())

	// Auth
	router.POST("/login", handlers.AuthPostLogin)
	router.POST("/user", handlers.AuthPostUser)
	authGroup.GET("/check", handlers.AuthGetCheckToken)

	// User info
	router.GET("/user/:id", handlers.UserInfoGetUserById)
	authGroup.PATCH("/user/:id", handlers.UserInfoPatchUserById)
	authGroup.DELETE("/user/:id", handlers.UserInfoDeleteUserById)

	log.Debugln("Created gin routing")

	port := fmt.Sprint(env.GetServerHttpPort())
	log.Debugln("Port from environment: " + port)
	router.Run(":" + port)

	log.Info("Server started on port " + port)
}
