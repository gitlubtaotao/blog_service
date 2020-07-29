package main

import (
	"blog_service/global"
	"blog_service/internal/model"
	"blog_service/internal/routes"
	"blog_service/pkg/logger"
	"blog_service/pkg/setting"
	"blog_service/pkg/tracer"
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"
)

// @title 博客系统
// @version 1.0
// @description Go 编程之旅, 一起用Go做项目
// @termsOfService https://github.com/gitlubtaotao/blog_service.git
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host localhost:8080
// @BasePath /api/v1
func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routes.NewRouter()
	s := http.Server{
		Addr:           global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	_ = s.ListenAndServe()
}

func init() {
	if err := setupSetting(); err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	if err := setupDBEngine(); err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
	if err := setupTracer(); err != nil {
		log.Fatalf("init.setupTracer err: %v", err)
	}
	if err := setupLogger(); err != nil {
		log.Fatalf("init.setupLogger err : %v", err)
	}
}


//读取系统配置文件
func setupSetting() error {
	set, err := setting.NewSetting()
	if err != nil {
		return err
	}
	if err = set.ReadSection("Server", &global.ServerSetting); err != nil {
		return err
	}
	if err = set.ReadSection("App", &global.AppSetting); err != nil {
		return err
	}
	if err = set.ReadSection("Database", &global.DatabaseSetting); err != nil {
		return err
	}
	if err = set.ReadSection("JWT", &global.JWTSetting); err != nil {
		return err
	}
	if err = set.ReadSection("Email", &global.EmailSetting); err != nil {
		return err
	}
	global.JWTSetting.Expire *= time.Second
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	global.AppSetting.DefaultContextTime *= time.Second
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename: global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:  600,
		MaxAge:   10,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}

func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer("blog-services", "127.0.0.1:6831")
	if err != nil {
		return err
	}
	global.Tracer = jaegerTracer
	return nil
}
