package V1

import (
	"blog_service/internal/service"
	"blog_service/pkg/app"
	"blog_service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func GetAuth(c *gin.Context) {
	param := service.AuthRequest{}
	response := app.NewResponse(c)
	if _, errs := app.BindAndValid(c, param); errs != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.NewAuthService()
	if err := svc.CheckAuth(&param); err != nil {
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist.WithDetails(err.Error()))
		return
	}
	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate.WithDetails(err.Error()))
		return
	}
	response.ToResponse(gin.H{"token": token})
}
