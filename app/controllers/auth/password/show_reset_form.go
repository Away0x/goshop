package password

import (
	"echo_shop/app/context"
	"echo_shop/app/models"
	"echo_shop/pkg/errno"
	"echo_shop/pkg/serializer"
)

// ShowResetForm 展示重置密码的页面
func ShowResetForm(c *context.AppContext) error {
	token := c.Param("token")

	pwd := new(models.PasswordReset)
	if err := models.Where("token = ?", token).First(&pwd).Error; err != nil {
		return errno.NotFoundErr.HTML()
	}

	return c.RenderHTML("auth/password/reset", serializer.Data{
		"token": token,
		"email": pwd.Email,
	})
}
