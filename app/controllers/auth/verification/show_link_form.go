package verification

import (
	"echo_shop/app/context"
)

// ShowLinkForm 展示发送激活用户链接邮件的页面
func ShowLinkForm(c *context.AppContext) error {
	return c.RenderHTML("auth/verification/show")
}
