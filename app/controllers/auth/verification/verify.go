package verification

import (
	"echo_shop/app/context"
	"echo_shop/app/models"
	"echo_shop/pkg/errno"
	"time"
)

// Verify 激活用户
func Verify(c *context.AppContext) error {
	user := new(models.User)
	if err := models.WhereFirst(&user, "activation_token = ?", c.Param("token")); err != nil {
		return errno.NotFoundErr.HTML()
	}

	// 更新用户
	user.Activated = models.TrueTinyint
	user.ActivationToken = ""
	now := time.Now()
	user.EmailVerifiedAt = &now
	if err := models.Save(user); err != nil {
		c.FlashDangerMessage(c.WEM(err, "用户激活失败"))
		return c.RedirectByName("verification.verify")
	}

	c.SessionLogin(user)
	c.FlashSuccessMessage("邮箱验证成功 ^_^")
	return c.RedirectToHomePage()
}
