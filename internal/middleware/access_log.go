package middleware

import (
	"blog_service/global"
	"blog_service/pkg/logger"
	"bytes"
	"github.com/gin-gonic/gin"
	"time"
)

// 记录每次一次请求的请求方法，方法调用开始时间,方法调用结束时间,方法响应结果
// 方法响应结果的状态码

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (a AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := a.body.Write(p); err != nil {
		return n, err
	}
	return a.ResponseWriter.Write(p)
}
func AccessLog() gin.HandlerFunc {
	return func(context *gin.Context) {
		bodyWriter := AccessLogWriter{
			ResponseWriter: context.Writer,
			body:           bytes.NewBufferString(""),
		}
		context.Writer = bodyWriter
		beginTime := time.Now().Unix()
		context.Next()
		endTime := time.Now().Unix()
		fields := logger.Fields{
			"request": context.Request.PostForm.Encode(),
			"response": bodyWriter.body.String(),
		}
		global.Logger.WithFields(fields).InfoF("access log: method: %s, status_code: %d," +
			"begin_time: %d, end_time: %d",context.Request.Method,bodyWriter.Status(),beginTime,endTime)
	}
}
