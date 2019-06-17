/// 登录
package login

import (
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	return c.String(200, "login")
}
