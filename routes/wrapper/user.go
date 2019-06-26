package wrapper

import (
	"echo_shop/app/context"
	"echo_shop/app/models"
)

// User : handler 中注入 user model
func User(handler func(*context.AppContext, *models.User) error) context.AppHandlerFunc {
	return func(c *context.AppContext) error {
		if currentUser, ok := c.CurrentUser(); ok {
			return handler(c, currentUser)
		}

		return c.RedirectToLoginPage()
	}
}
