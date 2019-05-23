package main

import (
	"echo_shop/bootstrap"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Debug = true
	bootstrap.SetupEcho(e)

	e.Logger.Fatal(e.Start(":1323"))
}
