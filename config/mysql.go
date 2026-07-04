package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewMysql() *gorm.DB {
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "13306")
	database := getEnv("DB_DATABASE", "simple_blog")
	user := getEnv("DB_USERNAME", "root")
	password := getEnv("DB_PASSWORD", "admin123")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		user, password, host, port, database,
	)

	logLevel := logger.Info

	// if !getEnvBool("APP_DEBUG", false) {
	// 	logLevel = logger.Silent
	// }

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})

	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %s", err.Error()))
	}

	return db
}
