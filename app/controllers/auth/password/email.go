package password

import (
	"echo_shop/app/context"
	"echo_shop/app/helpers/mail"
	"echo_shop/app/models"
	"echo_shop/app/request"

	"github.com/Away0x/validate"
)

type passwordEmailForm struct {
	validate.Base
	Email string
}

// Email 发送重置密码链接的 email
func Email(c *context.AppContext) error {
	req := &passwordEmailForm{
		Email: c.Param("email"),
	}

	if err, ok := validate.RunWithConfig(req, validate.Config{
		Plugins: validate.Plugins{
			request.EmailPlugin(req.Email),
		},
	}); !ok {
		c.ErrorFlash(err)
		return c.RedirectByName("password.show_link_form")
	}

	pwd := &models.PasswordReset{Email: req.Email}
	if err := models.Create(pwd); err != nil {
		c.FlashDangerMessage(c.WrapErrorMessage(err, "重置密码链接生成失败"))
		return c.RedirectByName("password.show_link_form")
	}

	// 发送邮件
	if err := mail.SendPasswordResetEmail(pwd); err != nil {
		c.FlashDangerMessage(c.WrapErrorMessage(err, "重置密码邮件发送失败"))
		models.Delete(pwd)
	} else {
		c.FlashSuccessMessage("重置密码已发送到你的邮箱上，请注意查收。")
	}

	return c.RedirectByName("password.show_link_form")
}
