/// 展示登录页面
package login

import (
	"echo_shop/pkg/errno"

	"github.com/labstack/echo/v4"
)

func Show(c echo.Context) error {
	// return c.String(200, "login.show")
	return errno.NotImplemented.RenderNoContent()
}
