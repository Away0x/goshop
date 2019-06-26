package register

import (
	"echo_shop/app/context"
	"echo_shop/app/models"
	"echo_shop/app/request"
	"errors"

	"github.com/Away0x/validate"
)

type registerForm struct {
	validate.BaseValidate
	Name                 string
	Email                string
	Password             string
	PasswordConfirmation string `json:"password_confirmation" form:"password_confirmation"`
	Captcha              string
	CaptchaID            string `json:"captcha_id" form:"captcha_id"`
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

	if err := user.Create(); err != nil {
		c.ErrorFlash(errors.New("用户创建失败: " + err.Error()))
		return c.RedirectToRegisterPage()
	}

	c.Login(user)
	c.FlashSuccessMessage("注册成功")
	return c.RedirectToHomePage()
}
