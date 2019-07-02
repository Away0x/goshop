package register

import (
	"echo_shop/app/context"
	"echo_shop/app/helpers/mail"
	"echo_shop/app/models"
	"echo_shop/app/request"
	"errors"

	"github.com/Away0x/validate"
)

type registerForm struct {
	validate.Base
	Name                 string
	Email                string
	Password             string
	PasswordConfirmation string `json:"password_confirmation" form:"password_confirmation"`
	Captcha              string
	CaptchaID            string `json:"captcha_id" form:"captcha_id"`
}

func (*registerForm) IsStrict() bool {
	return false
}

func (r *registerForm) Plugins() validate.Plugins {
	return validate.Plugins{
		request.UserNamePlugin(r.Name),
		request.EmailPlugin(r.Email),
		request.RegisterPasswordPlugin(r.Password, r.PasswordConfirmation),
		request.CaptchaPlugin(r.CaptchaID, r.Captcha),
	}
}

// Register 用户注册
func Register(c *context.AppContext) error {
	req := new(registerForm)

	if err := c.BindAndValidate(req); err != nil {
		c.ErrorFlash(err)
		return c.RedirectToRegisterPage()
	}

	// 创建用户
	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := models.Create(user); err != nil {
		c.ErrorFlash(errors.New("用户创建失败: " + err.Error()))
		return c.RedirectToRegisterPage()
	}

	c.Login(user)
	if err := mail.SendVerifyEmail(user); err != nil {
		c.FlashDangerMessage("邮件发送失败: " + err.Error())
	} else {
		c.FlashSuccessMessage("新的验证链接已发送到您的 E-mail")
	}

	return c.RedirectToHomePage()
}
