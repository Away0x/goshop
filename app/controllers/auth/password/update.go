package password

import (
	"echo_shop/app/context"
	"echo_shop/app/models"
	"echo_shop/app/request"

	"github.com/Away0x/validate"
)

type passwordUpdateForm struct {
	validate.Base
	Email                string
	Token                string
	Password             string
	PasswordConfirmation string `json:"password_confirmation" form:"password_confirmation"`
}

func (p *passwordUpdateForm) Plugins() validate.Plugins {
	return validate.Plugins{
		request.RegisterPasswordPlugin(p.Password, p.PasswordConfirmation),
	}
}

func (p *passwordUpdateForm) Validators() validate.Validators {
	return validate.Validators{
		"token": {
			validate.RequiredValidator(p.Token),
		},
	}
}

func (p *passwordUpdateForm) Messages() validate.Messages {
	return validate.Messages{
		"token": {
			"token 不能为空",
			"该 token 不存在",
		},
	}
}

// Update 重置密码
func Update(c *context.AppContext) error {
	token := c.FormValue("token")
	req := new(passwordUpdateForm)

	if err := c.BindAndValidate(req); err != nil {
		c.ErrorFlash(err)
		return c.RedirectByName("password.show_reset_form", token)
	}

	// update user
	user := new(models.User)
	if err := models.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.FlashDangerMessage(c.WrapErrorMessage(err, "该邮箱还未注册过用户"))
		return c.RedirectByName("password.show_reset_form", token)
	}
	user.Password = req.Password
	if err := models.Update(&user); err != nil {
		c.FlashDangerMessage(c.WrapErrorMessage(err, "重置密码失败"))
		return c.RedirectByName("password.show_reset_form", token)
	}

	// 删除 password reset token
	pwd := &models.PasswordReset{
		Email: req.Email,
		Token: req.Token,
	}
	if err := models.Delete(pwd); err != nil {
		c.FlashDangerMessage(c.WrapErrorMessage(err, "重置密码失败"))
		return c.RedirectByName("password.show_reset_form", token)
	}

	c.Login(user)
	c.FlashSuccessMessage("密码更新成功，您已成功登录！")
	return c.RedirectToHomePage()
}
