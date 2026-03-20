package album

import (
	"fmt"
	"net/http"
	"server4/internal/auth"
	"server4/internal/user"

	"github.com/gin-gonic/gin"
)

var lastID = 3

func GenerateAlbumID() string {
	lastID++
	return fmt.Sprintf("%d", lastID)
}

func EnSureArtistOwnerID(ctx *gin.Context, a *album) bool {
	role := auth.GetRole(ctx)
	userID := auth.GetUserID(ctx)

	if role != string(user.RoleArtist) {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "Only Artists could access"})
		return false
	}

	if a.OwnerID != userID {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "It is not belong to you"})
		return false
	}

	return true
}

func DeleteAlbumHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	existing := GetOne(id)

	if existing == nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "not found"})
		return
	}

	if !EnSureArtistOwnerID(ctx, existing) {
		return
	}

	if !DeleteOne(id) {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "can not delete album"})
		return
	}
	ctx.Status(http.StatusNoContent)
}
