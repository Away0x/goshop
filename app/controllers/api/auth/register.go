package auth

import (
	"echo_shop/app/context"
	"echo_shop/app/models"
	"echo_shop/app/request"
)

// Register 用户注册
func Register(c *context.AppContext) error {
	req := new(request.RegisterForm)
	req.APIForm = true // 不检查验证码

	if err := c.BindAndValidate(req); err != nil {
		return err
	}

	user := new(models.User)
	if err := models.WhereFirst(&user, "email = ?", req.Email); err == nil {
		return context.M{
			"email": {"该邮箱已经被使用了"},
		}
	}

	// 创建用户
	user = &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	if err := user.Create(); err != nil {
		return err
	}

	return c.RenderJSON(user)
}
