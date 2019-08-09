package echoinit

import (
	"echo_shop/bootstrap/config"
	"echo_shop/routes"
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/labstack/echo/v4"
)

// SetupRoute 路由配置
func SetupRoute(e *echo.Echo) {
	routes.Register(e)

	// 输出路由
	printRoutes(e)
}

func printRoutes(e *echo.Echo) {
	routes := make([]*echo.Route, 0)
	for _, item := range e.Routes() {
		if strings.HasPrefix(item.Name, "github.com") {
			continue
		}

		routes = append(routes, item)
	}

	routesStr, _ := json.MarshalIndent(struct {
		Count  int           `json:"count"`
		Routes []*echo.Route `json:"routes"`
	}{
		Count:  len(routes),
		Routes: routes,
	}, "", "  ")
	ioutil.WriteFile(config.String("APP.TEMP_DIR")+"/routes.json", routesStr, 0644)
}
