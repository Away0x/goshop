package bootstrap

import (
	"echo_shop/bootstrap/app"
	"echo_shop/bootstrap/config"
	"echo_shop/bootstrap/echoinit"
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
	e.HideBanner = true

	echoinit.SetupRoute(e)  // 配置路由
	echoinit.SetupRender(e) // 配置 render

	return e
}

// RunEcho -
func RunEcho() {
	e := SetupEcho()
	app.SaveApplication(e)

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
			e.Logger.Fatal(e.StartTLS(config.String("TLS.PORT"), certFile, keyFile))
		}()
	}

	// 启动 app
	e.Logger.Fatal(e.Start(config.String("APP.PORT")))

	// graceful -----------------------------------------------
	// if config.IsDev() {
	// 	e.Logger.Fatal(e.Start(config.String("APP.PORT")))
	// } else {
	// 	e.Server.Addr = config.String("APP.PORT")
	// 	e.Logger.Fatal(gracehttp.Serve(e.Server))
	// }
}

func pingServer() error {
	for i := 0; i < 5; i++ {
		resp, err := http.Get(config.String("APP.URL") + app.Application.Reverse("sd.health"))
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		log.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}

	return errors.New("Cannot connect to the router")
}
