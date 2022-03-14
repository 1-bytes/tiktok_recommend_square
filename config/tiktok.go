package config

import "tiktok/pkg/config"

func init() {
	config.Add("tiktok", config.StrMap{
		// 应用名称
		"api":    config.Env("TIKTOK_API"),
		"body":   config.Env("TIKTOK_BODY"),
		"cookie": config.Env("TIKTOK_COOKIE"),
	})
}
