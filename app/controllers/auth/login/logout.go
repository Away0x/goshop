/// 登出
package login

import (
	"echo_shop/app/context"
)

func Logout(c *context.AppContext) error {
	return c.String(200, "logout")
}
