package global

import (
	"blog_service/pkg/logger"
	"blog_service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettings
	Logger          *logger.Logger
)
