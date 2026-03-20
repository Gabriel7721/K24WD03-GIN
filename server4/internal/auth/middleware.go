package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	ContextUserIDKey   = "userID"
	ContextUsernameKey = "username"
	ContextRoleKey     = "role"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "Token not found"})
			return
		}
		// "Bearer kajlsdhfasdlkfha"
		// SplitN => ["Bearer", "kajlsdhfasdlkfha"]
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "Token not contain Bearer"})
			return
		}

		tokenStr := parts[1]
		claims, err := ParseToken(tokenStr)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "Token is expired"})
			return
		}

		ctx.Set(ContextUserIDKey, claims.ID)
		ctx.Set(ContextUsernameKey, claims.Username)
		ctx.Set(ContextRoleKey, claims.Role)

		ctx.Next()
	}
}
func RequireRole(role string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		value, isExisted := ctx.Get(ContextRoleKey)
		if !isExisted {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"msg": "Role Not Found"})
			return
		}
		currentRole := value.(string)
		if currentRole != role {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"msg": "You are not allowed to access this resource"})
			return
		}
		ctx.Next()
	}
}
