package useraddress

import (
	"echo_shop/app/context"
	"echo_shop/app/models"
	"echo_shop/pkg/errno"
)

// Destroy 删除收货地址
func Destroy(c *context.AppContext, user *models.User, tokenStr string) error {
	address := new(models.UserAddress)
	if err := c.ModelByID("id", &address); err != nil {
		return err
	}

	// 权限判断
	if err := c.AuthCheck(user.ID == address.UserID); err != nil {
		return err
	}

	if err := models.Delete(address); err != nil {
		return errno.DatabaseErr.SetErrorContent(err)
	}

	return c.RenderJSON(context.G{
		"id": address.ID,
	})
}
