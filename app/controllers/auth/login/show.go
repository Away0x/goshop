/// 展示登录页面
package login

import (
	"github.com/labstack/echo/v4"
)

func Show(c echo.Context) error {
	return c.String(200, "login.show")
}
