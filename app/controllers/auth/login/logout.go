/// 登出
package login

import (
	"github.com/labstack/echo/v4"
)

func Logout(c echo.Context) error {
	return c.String(200, "logout")
}
