package middleware

import (
	"echo_shop/config"

	session "echo_shop/pkg/session/echo-session"

	"github.com/labstack/echo/v4"
)

const (
	// SessionName -
	SessionName = "shop_session"
)

// Session -
func Session() echo.MiddlewareFunc {
	store := session.NewCookieStore([]byte(config.String("APP.KEY")))
	store.Options(session.Options{
		HttpOnly: true,
		Path:     "/",
		MaxAge:   86400 * 30,
	})

	return session.Sessions(SessionName, store)
}
