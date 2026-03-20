package main

import (
	"server4/internal/album"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.DELETE("/albums/:id", album.DeleteAlbumHandler)

	router.Run(":9999")
}
