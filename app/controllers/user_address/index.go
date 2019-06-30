package user_address

import (
	"echo_shop/app/context"
	"echo_shop/app/models"
)

func Index(c *context.AppContext, u *models.User) error {
	return c.RenderHTML("user_address/index", map[string]interface{}{})
}
