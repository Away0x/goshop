package echoinit

import (
	"echo_shop/bootstrap/config"
	"echo_shop/pkg/echohelper"
	"echo_shop/routes"

	"github.com/labstack/echo/v4"
)

// SetupRoute 路由配置
func SetupRoute(e *echo.Echo) {
	routes.Register(e)

	// 输出路由
	echohelper.PrintRoutes(e, config.String("APP.TEMP_DIR")+"/routes.json")
}
