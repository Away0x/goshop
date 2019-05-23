package bootstrap

import (
	"echo_shop/bootstrap/echoinit"

	"github.com/labstack/echo/v4"
)

// SetupEcho -
func SetupEcho(e *echo.Echo) {
	echoinit.SetupRoute(e)
	echoinit.SetupRender(e)
	echoinit.SetupValidate(e)
	echoinit.SetupError(e)
}
