package echoinit

import (
	"echo_shop/routes"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
)

func SetupRoute(e *echo.Echo) {
	// 项目静态文件配置
	e.Static("/public", "public")
	e.File("/favicon.ico", "public/favicon.ico")

	routes.Register(e)

	// 打印路由
	for _, v := range e.Routes() {
		fmt.Fprintf(os.Stderr, "[Route] "+"%-7s %-50s --> %s\n", v.Method, v.Name, v.Path)
	}
}
