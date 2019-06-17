package constants

import (
	"github.com/labstack/echo/v4"
)

const (
	// csrf -------------------------------
	// CsrfContexntName csrf token 在 echo context 中的 key name
	CsrfContexntName = "csrf"
	// CsrfCookieName csrf token 在 cookie 中的 key name
	CsrfCookieName = "_csrf"
	// CsrfValueName csrf token 在 header/params/query 中的 key name
	CsrfValueName = echo.HeaderXCSRFToken

	// method 重写 -------------------------------
	// MethodOverrideFromFormKeyName form 中的 key name
	MethodOverrideFromFormKeyName = "_method"
)
