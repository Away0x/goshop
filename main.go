package main

import (
	"echo_shop/bootstrap"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	bootstrap.SetupEcho(e)

	e.Logger.Fatal(e.Start(":1323"))
}
