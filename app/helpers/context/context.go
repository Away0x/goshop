package context

import (
	"echo_shop/app/models"

	"github.com/labstack/echo/v4"
)

// CustomContext -
type CustomContext struct {
	echo.Context
}

func (c *CustomContext) getUserBySession() (user *models.User, ok bool) {

}
