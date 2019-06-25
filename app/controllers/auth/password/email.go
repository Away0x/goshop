/// 发送重置密码链接的 email
package password

import (
	"echo_shop/app/context"
)

func Email(c *context.AppContext) error {
	return c.String(200, "password.email")
}
