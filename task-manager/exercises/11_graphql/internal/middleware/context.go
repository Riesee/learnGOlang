package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
)

type contextKey string

const UserIDKey contextKey = "userID"

func GinContextToContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("userID")
		if !exists {
			c.Next()
			return
		}
		ctx := context.WithValue(c.Request.Context(), UserIDKey, userID)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func GetUserIDFromContext(ctx context.Context) uint {
	data := ctx.Value(UserIDKey)
	if data == nil {
		return 1
	}
	userID, ok := data.(uint)
	if !ok {
		return 1
	}
	return userID
}