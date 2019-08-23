package echohelper

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/labstack/echo/v4"
)

// PrintRoutes 输出路由配置
func PrintRoutes(e *echo.Echo, filename string) {
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
	ioutil.WriteFile(filename, routesStr, 0644)
}
