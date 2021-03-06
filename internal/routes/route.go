package routes

import (
	_ "blog_service/docs"
	"blog_service/global"
	"blog_service/internal/middleware"
	"blog_service/internal/routes/api"
	"blog_service/internal/routes/api/V1"
	"blog_service/pkg/limiter"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"time"
)

var methodLimiter = limiter.NewMethodLimiter().AddBucket(limiter.BucketRule{
	Key:          "/auth",
	FillInterval: time.Second,
	Capacity:     10,
	Quantum:      10,
})

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware.AccessLog(), middleware.Recovery())
	r.Use(middleware.Translations())
	r.Use(middleware.AppInfo())
	r.Use(middleware.Tracing())
	r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTime), middleware.RateLimiter(methodLimiter))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	tag := V1.NewTag()
	article := V1.NewArticle()
	r.StaticFS("/static", gin.Dir(global.AppSetting.UploadSavePath, true))
	r.GET("/auth", V1.GetAuth)
	apiV1 := r.Group("/api/v1").Use(middleware.JWT())
	{
		apiV1.POST("/api/upload/file", api.UploadFile)
		apiV1.POST("/tags", tag.Create)
		apiV1.DELETE("/tags/:id", tag.Delete)
		apiV1.PUT("/tags/:id", tag.Update)
		apiV1.PATCH("/tags/:id/state", tag.ChangeState)
		apiV1.GET("/tags/:id", tag.Get)
		apiV1.GET("/tags", tag.List)
		apiV1.POST("/articles", article.Create)
		apiV1.DELETE("/articles/:id", article.Delete)
		apiV1.PUT("/articles/:id", article.Update)
		apiV1.PATCH("/articles/:id/state", article.ChangeState)
		apiV1.GET("/articles", article.List)
		apiV1.GET("/articles/:id", article.Get)
	}
	return r
}
