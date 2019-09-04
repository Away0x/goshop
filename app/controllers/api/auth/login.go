package auth

import (
	"echo_shop/app/auth/jwt"
	"echo_shop/app/context"
	"echo_shop/app/models"
	"echo_shop/app/request"
)

// Login 登录
func Login(c *context.AppContext) error {
	req := new(request.LoginForm)

	if err := c.BindAndValidate(req); err != nil {
		return err
	}

	user := new(models.User)
	if err := models.WhereFirst(&user, "email = ?", req.Email); err != nil {
		return c.EM(err, "email", "该邮箱还未注册")
	}
	if err := user.Compare(req.Password); err != nil {
		return c.EM(err, "password", "密码错误")
	}

	// 签发 token
	tokenInfo, err := jwt.Sign(user)
	if err != nil {
		return err
	}

	return c.RenderJSON(context.G{
		"user":  user.Serialize(),
		"token": tokenInfo.Token,
	})
}
