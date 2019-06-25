package bootstrap

import (
	"echo_shop/bootstrap/echoinit"
	"echo_shop/config"
	"errors"
	"net/http"
	"time"

	"github.com/lexkong/log"

	// "github.com/facebookgo/grace/gracehttp"

	"github.com/labstack/echo/v4"
)

// SetupEcho -
func SetupEcho() *echo.Echo {
	e := echo.New()
	e.Debug = config.IsDev()

	specialHandlers := echoinit.SetupRoute(e) // 配置路由
	echoinit.SetupRender(e)                   // 配置 render
	echoinit.SetupBinder(e)                   // 配置 binder
	echoinit.SetupError(e, specialHandlers)   // 配置统一错误处理

	return e
}

// RunEcho -
func RunEcho() {
	e := SetupEcho()
	config.SaveApplication(e)

	// 在启动 HTTP 端口前 go 一个 pingServer 协程，启动 HTTP 端口后，
	// 该协程不断地 ping /sd/health 路径，如果失败次数超过一定次数，则终止 HTTP 服务器进程
	// go func() {
	// 	time.Sleep(5 * time.Second) // 等待 server 启动
	// 	if err := pingServer(); err != nil {
	// 		log.Fatal("The router has no response, or it might took too long to start up.", err)
	// 	}
	// 	log.Info("The router has been deployed successfully.")
	// }()

	// 如果提供了 TLS 证书和私钥则启动 HTTPS 端口
	certFile := config.String("TLS.CERT_FILE")
	keyFile := config.String("TLS.KEY_FILE")
	if certFile != "" && keyFile != "" {
		go func() {
			e.Logger.Fatal(e.StartTLS(config.String("TLS.ADDR"), certFile, keyFile))
		}()
	}

	// 启动 app
	e.Logger.Fatal(e.Start(config.String("APP.ADDR")))

	// graceful -----------------------------------------------
	// if config.IsDev() {
	// 	e.Logger.Fatal(e.Start(config.String("APP.ADDR")))
	// } else {
	// 	e.Server.Addr = config.String("APP.ADDR")
	// 	e.Logger.Fatal(gracehttp.Serve(e.Server))
	// }
}

func pingServer() error {
	for i := 0; i < 5; i++ {
		resp, err := http.Get(config.String("APP.URL") + config.Application.Reverse("sd.health"))
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		log.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}

	return errors.New("Cannot connect to the router")
}
