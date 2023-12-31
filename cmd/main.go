package main

import (
	"file_share_server/global"
	"file_share_server/initialize"
	"file_share_server/model"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	initialize.Init()
	logger := logger.New (
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config {
			SlowThreshold: time.Duration(global.Config.Database.SlowThreshold) * time.Millisecond, // 慢SQL阈值
			//LogLevel: logger.Info, // 级别
			LogLevel: logger.LogLevel(global.Config.Database.LogLevel),
			Colorful: global.Config.Database.Colorful, // 彩色
		},
	)

	db, err := gorm.Open(mysql.Open(global.Config.Database.DSN), &gorm.Config {
		Logger: logger,
	})
	if err != nil {
	  panic("failed to connect database")
	}
	
	// 迁移 schema
	db.AutoMigrate(
		&model.ShareFile{},
	)
}