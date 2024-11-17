package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

var ctx = context.Background()

func RateLimitMiddleware(rdb *redis.Client, limit int, duration time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		count, err := rdb.Get(ctx, ip).Int()
		if err != nil && err != redis.Nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			c.Abort()
			return
		}

		if count >= limit {
			c.JSON(http.StatusTooManyRequests, gin.H{"message": fmt.Sprintf("Limit exceeded. Try again in %v seconds", duration.Seconds())})
			c.Abort()
			return
		}

		err = rdb.Set(ctx, ip, count+1, duration).Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			c.Abort()
			return
		}

		c.Next()
	}
}
