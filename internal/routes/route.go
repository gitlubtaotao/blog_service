package routes

import (
	"blog_service/internal/routes/api/V1"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	tag := V1.NewTag()
	article := V1.NewArticle()
	apiV1 := r.Group("/api/v1")
	{
		
		apiV1.POST("/tags", tag.Create)
		apiV1.DELETE("/tags/:id", tag.Delete)
		apiV1.PUT("/tags/:id", tag.Update)
		apiV1.PATCH("/tags/:id/state", tag.ChangeState)
		apiV1.GET("/tags/:id", tag.Get)
		apiV1.GET("/tags",tag.List)
		
		apiV1.POST("/articles", article.Create)
		apiV1.DELETE("/articles/:id", article.Delete)
		apiV1.PUT("/articles/:id", article.Update)
		apiV1.PATCH("/articles/:id/state", article.ChangeState)
		apiV1.GET("/articles", article.List)
		apiV1.GET("/articles/:id",article.Get)
	}
	return r
}
