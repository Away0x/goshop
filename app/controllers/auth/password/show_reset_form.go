package password

import (
	"echo_shop/app/context"
)

// ShowResetForm 展示重置密码的页面
func ShowResetForm(c *context.AppContext) error {
	return c.RenderHTML("auth/password/reset")
}
