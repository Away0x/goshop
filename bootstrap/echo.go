package bootstrap

import (
	"echo_shop/bootstrap/echoinit"

	"github.com/labstack/echo/v4"
)

// EchoEngine -
var EchoEngine *echo.Echo

// SetupEcho -
func SetupEcho() {
	e := echo.New()
	EchoEngine = e
	e.Debug = true

	echoinit.SetupRoute(e)
	echoinit.SetupRender(e)
	echoinit.SetupValidate(e)
	echoinit.SetupError(e)
}

// RunEcho -
func RunEcho() {
	EchoEngine.Logger.Fatal(EchoEngine.Start(":1323"))
}
