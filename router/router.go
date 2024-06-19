package router

import (
	"echo/common/logger"
	"echo/middlewares"
	"echo/router/admin"
	"echo/router/app"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"echo/router/common"

	"time"
)

func RunHttpServer(bing string) {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "ip=${remote_ip} time=${time_rfc3339}, method=${method}, uri=${uri}, status=${status}, latency_human=${latency_human}\n",
		Output: logger.EchoLog,
	}))
	e.Use(middleware.CORSWithConfig(
		middleware.CORSConfig{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{echo.POST, echo.GET, echo.OPTIONS, echo.PATCH, echo.DELETE},
			AllowCredentials: true,
			MaxAge:           int(time.Hour) * 24,
		}))
	e.Use(middlewares.RequestLog())
	e.Use(middleware.BodyDumpWithConfig(middlewares.DefaultBodyDumpConfig))

	registRouter := func(fn func(*echo.Group)) {
		if fn != nil {
			fn(e.Group(""))
		}
	}

	// ------------------ App Interface -------------------------
	// 登陆
	registRouter(func(eg *echo.Group) {
		g := eg.Group("/app")
		app.Register(g)
	})

	// ------------------ Company Platform Interface -------------------------

	// ------------------ Management Platform Interface -------------------------

	registRouter(func(eg *echo.Group) {
		common.Register(eg)
	})

	// 后台接口
	registRouter(func(eg *echo.Group) {
		g := eg.Group("/api/admin")
		g.Use(middlewares.AuthLogin())
		admin.Register(g)
	})

	e.Logger.Fatal(e.Start(bing))
}
