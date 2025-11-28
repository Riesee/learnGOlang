package middleware

import (
	"fmt"
	"fullstack/internal/cache"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func RateLimitMiddleware(maxRequests int, timeWindow time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetUint("userID")

		key := fmt.Sprintf("ratelimit:user:%d", userID)

		count, err := cache.Client.Incr(cache.Ctx, key).Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if count == 1 {
			cache.Client.Expire(cache.Ctx, key, timeWindow)
		}

		if count > int64(maxRequests) {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Rate limit exceeded",
				"rety_after": "60 seconds",
			})
			return
		}

        c.Header("X-RateLimit-Limit", fmt.Sprintf("%d", maxRequests))
        c.Header("X-RateLimit-Remaining", fmt.Sprintf("%d", maxRequests-int(count)))
		
		c.Next()
	}


}