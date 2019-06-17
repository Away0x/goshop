/// 展示注册页面
package register

import (
	"github.com/labstack/echo/v4"
)

func Show(c echo.Context) error {
	return c.String(200, "register.show")
}
