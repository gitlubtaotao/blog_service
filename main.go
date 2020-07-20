package main

import (
	"blog_service/global"
	"blog_service/internal/model"
	"blog_service/internal/routes"
	"blog_service/pkg/logger"
	"blog_service/pkg/setting"
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
func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routes.NewRouter()
	s := http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
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
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
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
