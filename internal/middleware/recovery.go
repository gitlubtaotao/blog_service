package middleware

import (
	"blog_service/global"
	"blog_service/pkg/app"
	"blog_service/pkg/email"
	"blog_service/pkg/errcode"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Recovery() gin.HandlerFunc {
	defaultMailer := email.NewEmail(&email.SMIPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		IsSSL:    global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.UserName,
		Password: global.EmailSetting.Password,
		From:     global.EmailSetting.From,
	})
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				global.Logger.WithCallersFrames().ErrorF(c,"panic error err: %v", err)
				errs := defaultMailer.SendMail(global.EmailSetting.To,
					fmt.Sprintf("异常抛出,发生时间: %d", time.Now().Unix()),
					fmt.Sprintf("错误信息： %v", err),
				)
				if errs != nil {
					global.Logger.PanicF(c,"Mail.sendMail is err: %v", errs)
				}
				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.Abort()
			}
			c.Next()
		}()
	}
}
