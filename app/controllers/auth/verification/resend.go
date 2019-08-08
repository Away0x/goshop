package verification

import (
	"echo_shop/app/context"
	"echo_shop/app/helpers/mail"
	"echo_shop/app/models"
)

// Resend 重新发送激活邮件
func Resend(c *context.AppContext, u *models.User) error {
	if u.IsActivated() {
		return c.RedirectBack("root")
	}

	if err := mail.SendVerifyEmail(u); err != nil {
		c.FlashDangerMessage(c.WrapErrorMessage(err, "邮件发送失败"))
	} else {
		c.FlashSuccessMessage("新的验证链接已发送到您的 E-mail")
	}

	return c.String(200, "verification.resend")
}
