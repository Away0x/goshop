package bootstrap

const (
	// 配置文件路径
	defaultConfigFilePath = "./config.yaml"
	// 配置文件格式
	defaultConfigFileType = "yaml"
)

// Run 启动项目
func Run() {
	RunWithConfig(defaultConfigFilePath, defaultConfigFileType)
}

// RunWithConfig 启动项目
func RunWithConfig(configFilePath, configFileType string) {
	// init config
	SetupConfig(configFilePath, configFileType)

	// init db
	db, _ := SetupDB()
	defer db.Close()

	// run echo
	RunEcho()
}
