package bootstrap

import (
	"echo_shop/bootstrap/echoinit"
	"echo_shop/config"

	// "github.com/facebookgo/grace/gracehttp"

	"github.com/labstack/echo/v4"
)

// EchoEngine -
var EchoEngine *echo.Echo

// SetupEcho -
func SetupEcho() {
	e := echo.New()
	EchoEngine = e
	e.Debug = config.IsDev()

	echoinit.SetupRoute(e)
	echoinit.SetupRender(e)
	echoinit.SetupValidate(e)
	echoinit.SetupError(e)
}

// RunEcho -
func RunEcho() {
	SetupEcho()
	EchoEngine.Logger.Fatal(EchoEngine.Start(config.String("APP.ADDR")))

	// graceful
	// if config.IsDev() {
	// 	EchoEngine.Logger.Fatal(EchoEngine.Start(config.String("APP.ADDR")))
	// } else {
	// 	EchoEngine.Server.Addr = config.String("APP.ADDR")
	// 	EchoEngine.Logger.Fatal(gracehttp.Serve(EchoEngine.Server))
	// }
}
