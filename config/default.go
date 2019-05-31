package config

import (
	"github.com/spf13/viper"
)

// 默认配置
var defaultConfigMap = map[string]interface{}{
	// app
	"APP.NAME":        "app",
	"APP.RUNMODE":     "release",
	"APP.ADDR":        ":9000",
	"APP.KEY":         "Rtg8BPKNEf2mB4mgvKONGPZZQSaJWNLijxR42qRgq0iBb5",
	"APP.ENABLE_CSRF": true,

	// db
	"DB.CONNECTION": "mysql",
	"DB.HOST":       "127.0.0.1",
	"DB.PORT":       "3306",
	"DB.DATABASE":   viper.GetString("APP.NAME"),
	"DB.USERNAME":   "root",
	"DB.PASSWORD":   "",
}

// 设置配置默认值
func setupDefaultConfig() {
	for k, v := range defaultConfigMap {
		viper.SetDefault(k, v)
	}
}
