/// 重置密码
package password

import (
	"echo_shop/app/context"
)

func Update(c *context.AppContext) error {
	return c.String(200, "password.update")
}
