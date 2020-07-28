package middleware

import (
	"blog_service/pkg/app"
	"blog_service/pkg/errcode"
	"blog_service/pkg/limiter"
	"github.com/gin-gonic/gin"
)

func RateLimiter(l limiter.IFace) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
			count := bucket.TakeAvailable(1)
			if count == 0 {
				response := app.NewResponse(c)
				response.ToErrorResponse(errcode.TooManyRequests)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
