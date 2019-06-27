package bootstrap

// Run 启动项目
func Run() {
	// init db
	db, _ := SetupDB()
	defer db.Close()

	// run echo
	RunEcho()
}
