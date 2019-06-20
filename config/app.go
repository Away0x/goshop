package config

import (
	"github.com/labstack/echo/v4"
)

var (
	// Application echo Engine
	Application *echo.Echo
)

// SaveApplication 存储 echo engine
func SaveApplication(e *echo.Echo) {
	Application = e
}
