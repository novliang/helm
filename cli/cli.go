package main

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/novliang/helm/app"
	"github.com/novliang/helm/engine"
	"github.com/novliang/helm/global"
	"github.com/novliang/helm/initialization"
	"time"
)

func main() {
	initialization.Config()
	initialization.Mysql()
	initialization.Logger()
	e := engine.NewEcho(engine.EchoConfig{
		Middleware: []echo.MiddlewareFunc{
			middleware.Logger(),
			middleware.Recover(),
		},
		Logger: global.HELM_LOGGER,
	})
	app.Router()(e.Echo)
	address := fmt.Sprintf(":%d", global.HELM_CONF.System.Addr)
	s := endless.NewServer(address, e)
	s.ReadHeaderTimeout = 10 * time.Millisecond
	s.WriteTimeout = 10 * time.Second
	s.MaxHeaderBytes = 1 << 20
	global.HELM_LOGGER.Error(s.ListenAndServe())
}
