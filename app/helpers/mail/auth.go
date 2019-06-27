package mail

import (
	"echo_shop/app/models"
	"echo_shop/config"
)

// SendVerifyEmail 发送激活用户的邮件
func SendVerifyEmail(u *models.User) error {
	subject := "感谢注册 Weibo 应用！请确认你的邮箱。"
	tpl := "mail/verify.html"
	verifyURL := config.Application.Reverse("verification.verify", u.ActivationToken)

	return SendMail([]string{u.Email}, subject, tpl, map[string]interface{}{
		"url": config.String("APP.URL") + verifyURL,
	})
}
