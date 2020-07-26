package service

import (
	"blog_service/global"
	dao2 "blog_service/internal/dao"
	"errors"
)

type AuthRequest struct {
	AppKey    string `form:"app_key" binding:"required"`
	AppSecret string `form:"app_secret" binding:"required"`
}
type IAuthService interface {
	CheckAuth(param *AuthRequest) error
}

type AuthService struct {
	dao dao2.IAuthDao
}

func (a *AuthService) CheckAuth(param *AuthRequest) error {
	auth, err := a.dao.GetAuth(param.AppKey, param.AppSecret)
	if err != nil {
		return err
	}
	if auth.ID > 0 {
		return nil
	}
	return errors.New("auth info does not exist")
}

func NewAuthService() IAuthService {
	return &AuthService{dao: dao2.NewAuthDao(global.DBEngine)}
}
