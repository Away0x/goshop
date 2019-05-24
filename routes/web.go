package routes

import (
	"echo_shop/app/controllers/page"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func registerWeb(e *echo.Echo, middlewares ...echo.MiddlewareFunc) {
	ee := e.Group("")

	// session
	ee.Use(session.MiddlewareWithConfig(session.Config{
		Store: sessions.NewCookieStore([]byte("secret")),
	}))

	// csrf
	ee.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		Skipper:        middleware.DefaultSkipper,
		TokenLength:    32,
		TokenLookup:    "header:" + echo.HeaderXCSRFToken,
		ContextKey:     "csrf",
		CookieName:     "_csrf",
		CookiePath:     "/",
		CookieHTTPOnly: true,
		CookieMaxAge:   86400,
	}))

	ee.GET("/", page.Root).Name = "root"
}
