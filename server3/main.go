package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/albums", GetAlbums)
	r.POST("/albums", PostAlbum)
	r.GET("/albums/:id", GetOne)
	r.PUT("/albums/:id", PutOne)
	r.PATCH("/albums/:id", PatchOne)
	r.DELETE("/albums/:id", DeleteOne)

	r.Run(":9999")
}

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}
type albumForPatch struct {
	Title  *string  `json:"title"`
	Artist *string  `json:"artist"`
	Price  *float64 `json:"price"`
}

var albums = []album{
	{"1", "Thriller", "Michael Jackson", 1500.99},
	{"2", "Back in Black", "AC/DC", 2500.99},
	{"3", "The Bodyguard", "Whitney Houston", 3500.99},
}

func GetAlbums(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, albums)
}
func PostAlbum(ctx *gin.Context) {
	var a album
	ctx.BindJSON(&a)
	a.ID = strconv.Itoa(len(albums) + 1)
	albums = append(albums, a)
	ctx.JSON(http.StatusCreated, a)
}
func GetOne(ctx *gin.Context) {
	id := ctx.Param("id")
	for _, a := range albums {
		if a.ID == id {
			ctx.JSON(http.StatusOK, a)
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"id": id, "msg": "Album Not Found"})
}
func PutOne(ctx *gin.Context) {
	id := ctx.Param("id")
	var a album
	ctx.BindJSON(&a)
	a.ID = id
	for i := range albums {
		if albums[i].ID == id {
			albums[i] = a
			ctx.JSON(http.StatusOK, a)
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"id": id, "msg": "Album Not Found"})
}
func PatchOne(ctx *gin.Context) {
	id := ctx.Param("id")
	var a albumForPatch
	ctx.BindJSON(&a)

	for i := range albums {
		if albums[i].ID == id {
			if a.Title != nil {
				albums[i].Title = *a.Title
			}
			if a.Artist != nil {
				albums[i].Artist = *a.Artist
			}
			if a.Price != nil {
				albums[i].Price = *a.Price
			}

			ctx.JSON(http.StatusOK, albums[i])
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"id": id, "msg": "Album Not Found"})
}

func DeleteOne(ctx *gin.Context) {
	id := ctx.Param("id")
	for i := range albums {
		if albums[i].ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			ctx.Status(http.StatusNoContent)
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"id": id, "msg": "Album Not Found"})
}
