package V1

import (
	"blog_service/internal/service"
	"blog_service/pkg/app"
	"blog_service/pkg/convert"
	"blog_service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Article struct {
}

// @Summary 获取单个文章
// @Produce json
// @Param id body string true "文章对应的ID"
// @Success 200 { object } model.TagSwagger  <-- This is a user defined struct "成功"
// @Failure 400 { object } errcode.Error  <-- This is a user defined struct "请求错误"
// @Failure 500 { object } errcode.Error  <-- This is a user defined struct "内部错误"
// @Router /api/v1/tags [get]
func (t *Article) Get(c *gin.Context) {
	param := service.ArticleRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	_, errs := app.BindAndValid(c, param)
	if errs != nil {
		app.NewResponse(c).ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.NewArticleService(c)
	article, err := svc.GetArticle(&param)
	if err != nil {
		app.NewResponse(c).ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}
	response.ToResponse(article)
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
