package cmd

import (
	"echo_shop/bootstrap/config"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	// 默认配置文件路径
	defaultConfigFilePath = "config.yaml"
	// 配置文件格式
	configFileType = "yaml"
)

var configFilePath string

var rootCmd = &cobra.Command{
	Use: "echo_shop",
}

// Execute -
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// 配置文件路径: --cofig config_path
	rootCmd.PersistentFlags().
		StringVar(&configFilePath, "config", "", "config file (default is $APP/"+defaultConfigFilePath+")")
}

func initConfig() {
	// setup config
	if configFilePath == "" {
		configFilePath = defaultConfigFilePath
	}

	config.InitConfig(configFilePath, configFileType)
}
