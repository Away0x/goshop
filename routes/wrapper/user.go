package wrapper

import (
	"echo_shop/app/context"
	"echo_shop/app/models"
	"echo_shop/pkg/constants"
	"echo_shop/pkg/errno"
)

// User : handler 中注入 user model
func User(handler func(*context.AppContext, *models.User) error) context.AppHandlerFunc {
	return func(c *context.AppContext) error {
		if currentUser := c.CurrentUser(); currentUser != nil {
			return handler(c, currentUser)
		}

		if constants.NeedResponseJSON(c) {
			return errno.LoginRequiredErr
		}
		return c.RedirectToLoginPage()
	}
}
