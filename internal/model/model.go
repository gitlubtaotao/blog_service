package model

import (
	"blog_service/global"
	"blog_service/pkg/database"
	"blog_service/pkg/setting"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

func NewDBEngine(setting *setting.DatabaseSettings) (*gorm.DB, error) {
	var dialect gorm.Dialector
	if setting.DBType == "mysql" {
		dialect = database.MysqlDBDialect(setting)
	}
	logLevel := logger.Silent
	if global.ServerSetting.RunMode == "debug" {
		logLevel = logger.Info
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logLevel,    // Log level
			Colorful:      false,       // Disable color
		},
	)
	db, err := gorm.Open(dialect, &gorm.Config{
		PrepareStmt: true,
		Logger:      newLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   setting.TablePrefix,
			SingularTable: false,
		},
	})
	if err != nil {
		return nil, err
	}
	if sqlDB, err := db.DB(); err != nil {
		return nil, err
	} else {
		sqlDB.SetMaxIdleConns(setting.MaxIdleConns)
		sqlDB.SetMaxOpenConns(setting.MaxOpenConns)
	}
	return db, nil
}
