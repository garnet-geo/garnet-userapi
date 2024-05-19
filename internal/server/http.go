package server

import (
	"fmt"
	"log"

	"github.com/garnet-geo/garnet-userapi/internal/env"
	"github.com/garnet-geo/garnet-userapi/internal/handlers"
	"github.com/gin-gonic/gin"
	// jwtware "github.com/appleboy/gin-jwt/v2"mm,
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

	port := fmt.Sprint(env.GetServerHttpPort())
	router.Run(":" + port)

	log.Default().Println("Server started on port " + port)
}
