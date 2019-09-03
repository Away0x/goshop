package wrapper

import (
	"echo_shop/app/auth/jwt"
	"echo_shop/app/context"
	"echo_shop/app/models"
	"echo_shop/pkg/constants"
	"echo_shop/pkg/errno"
)

// UserBySession : handler 中注入 user model
func UserBySession(handler func(*context.AppContext, *models.User) error) context.AppHandlerFunc {
	return func(c *context.AppContext) error {
		if currentUser := c.SessionCurrentUser(); currentUser != nil {
			return handler(c, currentUser)
		}

		if constants.IsAPIRequest(c) {
			return errno.LoginRequiredErr
		}
		return c.RedirectToLoginPage()
	}
}

// GetToken : 注入 token
func GetToken(handler func(*context.AppContext, string) error) context.AppHandlerFunc {
	return func(c *context.AppContext) error {
		tokenStr, err := jwt.GetToken(c)
		if err != nil {
			return err
		}

		return handler(c, tokenStr)
	}
}

// GetTokenAndUser : handler 中注入 token string 和 user
func GetTokenAndUser(handler func(*context.AppContext, *models.User, string) error) context.AppHandlerFunc {
	return func(c *context.AppContext) error {
		tokenStr, curUser, err := jwt.GetTokenAndUser(c)
		if err != nil {
			return err
		}

		return handler(c, curUser, tokenStr)
	}
}
