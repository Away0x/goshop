package password

import (
	"echo_shop/app/context"
)

// ShowLinkForm 展示发送重置密码链接 email 的页面
func ShowLinkForm(c *context.AppContext) error {
	return c.RenderHTML("auth/password/email")
}
