package config

import (
	"fmt"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
)

const (
	// RunmodeProduction 生产环境
	RunmodeProduction = "production"
	// RunmodeStaging 准生产环境
	RunmodeStaging = "staging"
	// RunmodeDevelopment 调试、开发环境
	RunmodeDevelopment = "development"
	// RunmodeTest 测试环境
	RunmodeTest = "test"
)

// InitConfig 初始化 config
func InitConfig(configFilePath, configFileType string) {
	// 初始化 viper 配置
	viper.SetConfigFile(configFilePath)
	viper.SetConfigType(configFileType)

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("读取配置文件失败，请检查 %s 配置文件是否存在: %v", configFilePath, err))
	}

	// 设置配置默认值
	setupDefaultConfig()

	// 环境变量 (设置环境变量: export ECHO_SHOP_APP_RUNMODE=development)
	viper.AutomaticEnv()
	viper.SetEnvPrefix(String("APP.NAME")) // 环境变量前缀
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

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
