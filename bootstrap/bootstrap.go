package bootstrap

import (
	_ "echo_shop/config" // config
)

// Run 启动项目
func Run() {
	db, _ := SetupDB()
	defer db.Close()

	RunEcho()
}
