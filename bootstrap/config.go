package bootstrap

import (
	"echo_shop/config"
)

// SetupConfig 初始化 config
func SetupConfig(configFilePath, configFileType string) {
	config.InitConfig(configFilePath, configFileType)
}
