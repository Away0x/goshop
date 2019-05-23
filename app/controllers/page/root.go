package page

import (
	"github.com/labstack/echo"
)

// Root -
func Root(c echo.Context) error {
	return c.Render(200, "root.html", map[string]interface{}{
		"name": "root",
	})
}
