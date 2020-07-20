package V1

import (
	"blog_service/pkg/app"
	"blog_service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Article struct {
}

func (t *Article) Get(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
}

func (t *Article) List(c *gin.Context) {

}

func (t *Article) Create(c *gin.Context) {

}

func (t *Article) Update(c *gin.Context) {

}

func (t *Article) Delete(c *gin.Context) {

}

func (t *Article) ChangeState(c *gin.Context) {

}
func NewArticle() Article {
	return Article{}
}
