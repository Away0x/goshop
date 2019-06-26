package login

import (
	"echo_shop/app/context"
)

// Logout 登出
func Logout(c *context.AppContext) error {
	c.Logout()
	c.FlashSuccessMessage("您已成功退出！")
	return c.RedirectToLoginPage()
}
