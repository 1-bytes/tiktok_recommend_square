package config

import "tiktok/pkg/config"

func init() {
	config.Add("tiktok", config.StrMap{
		// 应用名称
		"email":    config.Env("TIKTOK_EMAIL"),
		"password": config.Env("TIKTOK_PASSWORD"),
		"api":      config.Env("TIKTOK_API"),
		"body":     config.Env("TIKTOK_BODY"),
	})
}
