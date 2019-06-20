package middleware

import (
	"echo_shop/pkg/constants"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Csrf() echo.MiddlewareFunc {
	return middleware.CSRFWithConfig(middleware.CSRFConfig{
		Skipper:     middleware.DefaultSkipper,
		TokenLength: 32,
		// - "header:<name>"
		// - "form:<name>"
		// - "query:<name>"
		TokenLookup:    "form:" + constants.CsrfValueName, // FormValue 中需有 constants.CsrfValueName
		ContextKey:     constants.CsrfContexntName,
		CookieName:     constants.CsrfCookieName,
		CookiePath:     "/",
		CookieHTTPOnly: true,
		CookieMaxAge:   86400,
	})
}
