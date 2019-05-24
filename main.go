package main

import (
	"echo_shop/bootstrap"
	_ "echo_shop/config"
)

func main() {
	bootstrap.SetupEcho()
	bootstrap.RunEcho()
}
