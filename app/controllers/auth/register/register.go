package register

import (
	"echo_shop/app/context"
	"echo_shop/app/helpers/mail"
	"echo_shop/app/models"
	"echo_shop/app/request"
)

// Register 用户注册
func Register(c *context.AppContext) error {
	req := new(request.RegisterForm)

	if err := c.BindAndValidate(req); err != nil {
		c.ErrorFlash(err)
		return c.RedirectToRegisterPage()
	}

	user := new(models.User)
	if err := models.WhereFirst(&user, "email = ?", req.Email); err == nil {
		c.ErrorFlashWM("email", "该邮箱已经被使用了")
		return c.RedirectToRegisterPage()
	}

	// 创建用户
	user = &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := models.Create(user); err != nil {
		c.ErrorFlash(c.WrapError(err, "用户创建失败"))
		return c.RedirectToRegisterPage()
	}

	c.SessionLogin(user)
	if err := mail.SendVerifyEmail(user); err != nil {
		c.FlashDangerMessage(c.WrapErrorMessage(err, "邮件发送失败"))
	} else {
		c.FlashSuccessMessage("新的验证链接已发送到您的 E-mail")
	}

	return c.RedirectToHomePage()
}
