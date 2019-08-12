package mail

import (
	"echo_shop/app/models"
	"echo_shop/bootstrap/app"
	"echo_shop/bootstrap/config"
)

// SendVerifyEmail 发送激活用户的邮件
func SendVerifyEmail(u *models.User) error {
	subject := "感谢注册 Weibo 应用！请确认你的邮箱。"
	tpl := "mail/verify.html"
	verifyURL := app.Application.Reverse("verification.verify", u.ActivationToken)

	return SendMail([]string{u.Email}, subject, tpl, map[string]interface{}{
		"url": config.String("APP.URL") + verifyURL,
	})
}

// SendPasswordResetEmail 发送重置密码的邮件
func SendPasswordResetEmail(pwd *models.PasswordReset) error {
	subject := "重置密码！请确认你的邮箱。"
	tpl := "mail/reset_password.html"
	resetPasswordURL := app.Application.Reverse("password.show_reset_form", pwd.Token)

	return SendMail([]string{pwd.Email}, subject, tpl, map[string]interface{}{
		"url": config.String("APP.URL") + resetPasswordURL,
	})
}
