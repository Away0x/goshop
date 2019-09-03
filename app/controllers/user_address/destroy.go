package useraddress

import (
	"echo_shop/app/context"
	"echo_shop/app/models"
	"echo_shop/pkg/errno"
)

// Destroy 删除收货地址
func Destroy(c *context.AppContext, u *models.User) error {
	address := new(models.UserAddress)
	if err := c.ModelByID("user_address", &address); err != nil {
		return err.JSON()
	}

	// 权限判断
	if err := c.AuthCheck(u.ID == address.UserID); err != nil {
		return err.JSON()
	}

	if err := models.Delete(address); err != nil {
		return errno.DatabaseErr.SetErrorContent(err).JSON()
	}

	return c.RenderJSON(nil)
}
