package config

import (
	"fmt"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
)

const (
	// 配置文件路径
	configFilePath = "./config.yaml"
	// 配置文件格式
	configFileType = "yaml"

	// RunmodeProduction 生产模式
	RunmodeProduction = "production"
	// RunmodeDevelopment 调试、开发模式
	RunmodeDevelopment = "development"
	// RunmodeTest 测试模式
	RunmodeTest = "test"
)

func init() {
	// 初始化 viper 配置
	viper.SetConfigFile(configFilePath)
	viper.SetConfigType(configFileType)
	// 环境变量 (设置环境变量: export ECHO_SHOP_APP_RUNMODE=development)
	viper.AutomaticEnv()
	viper.SetEnvPrefix("ECHO_SHOP") // 环境变量前缀
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("读取配置文件失败，请检查 config.yaml 配置文件是否存在: %v", err))
	}

	// 设置配置默认值
	setupDefaultConfig()

	// 初始化日志配置
	initLog()

	// 监听配置文件变化
	watchConfig()
}

// 监控配置文件变化
func watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(ev fsnotify.Event) {
		log.Infof("Config file changed: %s", ev.Name)
	})
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

// IsDev 是否为开发模式
func IsDev() bool {
	return String("APP.RUNMODE") == RunmodeDevelopment
}
