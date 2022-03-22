package bootstrap

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"tiktok/pkg/config"
	"time"
)

var DB *gorm.DB

// SetupMySQL 初始化 MySQL.
func SetupMySQL() {
	var (
		host     = config.GetString("mysql.host")
		port     = config.GetString("mysql.port")
		username = config.GetString("mysql.username")
		password = config.GetString("mysql.password")
		db       = config.GetString("mysql.db")
	)
	MySQL := ConnectMySQL(host, port, username, password, db)
	sqlDB, _ := MySQL.DB()

	// 设置最大连接数
	sqlDB.SetMaxOpenConns(100)
	// 设置最大空闲连接
	sqlDB.SetMaxIdleConns(25)
	// 设置每个连接的过期时间
	sqlDB.SetConnMaxLifetime(5 * time.Minute)
}

// ConnectMySQL 连接 MySQL.
func ConnectMySQL(host string, port string, username string, password string, db string) *gorm.DB {
	cfg := mysql.New(mysql.Config{
		DSN: fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			username, password, host, port, db,
		),
	})
	logger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Warn, // Log level
			Colorful:      false,       // 禁用彩色打印
		},
	)

	DB, _ = gorm.Open(cfg, &gorm.Config{
		Logger: logger,
	})
	return DB
}
