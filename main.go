package main

import (
	"echo_shop/bootstrap"
	_ "echo_shop/config"
)

func main() {
	db, _ := bootstrap.SetupDB()
	defer db.Close()

	bootstrap.RunEcho()
}
