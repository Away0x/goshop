package bootstrap

import (
	"echo_shop/bootstrap/echoinit"
	"echo_shop/config"

	// "github.com/facebookgo/grace/gracehttp"

	"github.com/labstack/echo/v4"
)

// SetupEcho -
func SetupEcho() *echo.Echo {
	e := echo.New()
	e.Debug = config.IsDev()

	echoinit.SetupRoute(e)
	echoinit.SetupRender(e)
	echoinit.SetupValidate(e)
	echoinit.SetupError(e)

	return e
}

// RunEcho -
func RunEcho() {
	e := SetupEcho()
	config.SaveApplication(e)
	// 启动 app
	e.Logger.Fatal(e.Start(config.String("APP.ADDR")))

	// graceful
	// if config.IsDev() {
	// 	e.Logger.Fatal(e.Start(config.String("APP.ADDR")))
	// } else {
	// 	e.Server.Addr = config.String("APP.ADDR")
	// 	e.Logger.Fatal(gracehttp.Serve(e.Server))
	// }
}
