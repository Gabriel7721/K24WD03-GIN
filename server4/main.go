package main

import (
	"server4/internal/album"
	"server4/internal/auth"
	"server4/internal/user"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/login", user.LoginHandler)

	api := router.Group("/api")
	api.Use(auth.AuthMiddleware())

	// NOTE: Homeworks 1
	api.GET("/albums")
	api.GET("/albums/:id")
	api.POST("/albums")
	api.PUT("/albums/:id")
	api.PATCH("/albums/:id")
	// ====

	// NOTE: Homeworks 2
	// Kết nối với MySQL/MongoDB

	api.DELETE("/albums/:id", album.DeleteAlbumHandler)

	router.Run(":9999")
}
