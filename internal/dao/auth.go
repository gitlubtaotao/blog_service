package dao

import (
	"blog_service/internal/model"
	"gorm.io/gorm"
)

type IAuthDao interface {
	GetAuth(appKey string, appSecret string) (model.Auth, error)
}

type AuthDao struct {
	*Dao
}

func (a *AuthDao) GetAuth(appKey string, appSecret string) (model.Auth, error) {
	auth := model.Auth{AppSecret: appSecret, AppKey: appKey}
	return auth.Get(a.engine)
}

func NewAuthDao(db *gorm.DB) IAuthDao {
	return &AuthDao{NewDao(db)}
}
