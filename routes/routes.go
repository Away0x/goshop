package routes

import (
	"echo_shop/app/controllers/page"

	"github.com/labstack/echo"
)

// Register 注册路由
func Register(e *echo.Echo) {
	e.GET("/", page.Root).Name = "root"
}
