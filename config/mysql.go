package config

import "tiktok/pkg/config"

func init() {
	config.Add("mysql", config.StrMap{
		// 应用名称
		"host":     config.Env("MYSQL_HOST", "127.0.0.1"),
		"port":     config.Env("MYSQL_PORT", "3306"),
		"username": config.Env("MYSQL_USERNAME", "root"),
		"password": config.Env("MYSQL_PASSWORD"),
		"db":       config.Env("MYSQL_DB"),
		"table":    config.Env("MYSQL_TABLE"),
	})
}
