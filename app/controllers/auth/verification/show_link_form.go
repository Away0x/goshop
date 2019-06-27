package verification

import (
	"echo_shop/app/context"
	"echo_shop/app/models"
)

// ShowLinkForm 展示发送激活用户链接邮件的页面
func ShowLinkForm(c *context.AppContext, u *models.User) error {
	if u.IsActivated() {
		return c.RedirectBack("root")
	}

	return c.RenderHTML("auth/verification/show")
}
