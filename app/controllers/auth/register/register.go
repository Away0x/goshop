/// 注册
package register

import (
	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {
	return c.String(200, "register")
}
