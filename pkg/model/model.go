package model

import (
	"github.com/sebastiankennedy/go-web-skeleton/pkg/logger"
	"gorm.io/gorm"
	// GORM 的 调试功能，可以在命令行中查看请求的 SQL 信息
	gormlogger "gorm.io/gorm/logger"
	// GORM 的 MySQL 数据库驱动导入
	"gorm.io/driver/mysql"
)

// DB gorm.DB 对象
var DB *gorm.DB

// ConnectDataBase 初始化模型
func ConnectDataBase() *gorm.DB {

	var err error

	config := mysql.New(mysql.Config{
		DSN: "root:123456@tcp(127.0.0.1:33060)/example?charset=utf8&parseTime=True&loc=Local",
	})

	// 准备数据库连接池
	DB, err = gorm.Open(config, &gorm.Config{
		Logger: gormlogger.Default.LogMode(gormlogger.Info),
	})
	logger.Error(err)

	return DB
}
