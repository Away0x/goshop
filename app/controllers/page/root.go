package page

import (
	"echo_shop/app/context"
)

// Root -
func Root(c *context.AppContext) error {
	return c.RenderHTML("root")
}
