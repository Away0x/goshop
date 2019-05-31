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

	// RunmodeRelease 生产模式
	RunmodeRelease = "release"
	// RunmodeDebug 调试、开发模式
	RunmodeDebug = "debug"
	// RunmodeTest 测试模式
	RunmodeTest = "test"
)

func init() {
	// 初始化 viper 配置
	viper.SetConfigFile(configFilePath)
	viper.SetConfigType(configFileType)

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("读取配置文件失败，请检查 config.yaml 配置文件是否存在: %v", err))
	}

	// 设置配置默认值
	setupDefaultConfig()

	// 初始化日志配置
	initLog()
}

// String -
func String(key string) string {
	return viper.GetString(key)
}

// DefaultString -
func DefaultString(key string, defaultVal string) string {
	v := viper.GetString(key)
	if v == "" {
		return defaultVal
	}

	return v
}

// Int -
func Int(key string) int {
	return viper.GetInt(key)
}

// DefaultInt -
func DefaultInt(key string, defaultVal int) int {
	v := viper.GetInt(key)
	if v == 0 {
		return defaultVal
	}

	return v
}

// Bool -
func Bool(key string) bool {
	return viper.GetBool(key)
}
