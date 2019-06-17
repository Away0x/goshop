package middleware

import (
	"echo_shop/config"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// Session -
func Session() echo.MiddlewareFunc {
	store := sessions.NewCookieStore([]byte(config.String("APP.KEY")))
	store.Options = &sessions.Options{
		HttpOnly: true,
		Path:     "/",
		MaxAge:   86400 * 30,
	}

	return session.MiddlewareWithConfig(session.Config{
		Store: store,
	})
}
