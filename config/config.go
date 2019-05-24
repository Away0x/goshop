package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	// 配置文件路径
	configFilePath = "./config.yaml"
	// 配置文件格式
	configFileType = "yaml"
)

func init() {
	// 初始化 viper 配置
	viper.SetConfigFile(configFilePath)
	viper.SetConfigType(configFileType)

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("读取配置文件失败，请检查 config.yaml 配置文件是否存在: %v", err))
	}

	// 初始化日志配置
	initLog()
}
