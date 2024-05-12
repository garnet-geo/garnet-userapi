package server

import (
	"fmt"
	"log"

	"github.com/garnet-geo/garnet-userapi/internal/env"
	"github.com/garnet-geo/garnet-userapi/internal/handlers"
	"github.com/gin-gonic/gin"
)

func InitServer() {
	router := gin.Default()

	// Auth
	router.POST("/login", handlers.AuthPostLogin)
	router.POST("/user", handlers.AuthPostUser)
	router.GET("/check", handlers.AuthGetCheckToken)

	// User info
	router.GET("/user/:id", handlers.UserInfoGetUserById)
	router.PATCH("/user/:id", handlers.UserInfoPatchUserById)
	router.DELETE("/user/:id", handlers.UserInfoDeleteUserById)

	port := fmt.Sprint(env.GetServerHttpPort())
	router.Run(":" + port)

	log.Default().Println("Server started on port " + port)
}
