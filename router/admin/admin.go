package admin

import (
	"echo/middlewares"
	"echo/router/admin/adminlog"
	"echo/router/admin/menu"
	"echo/router/admin/no"
	"github.com/labstack/echo"

	"echo/router/admin/user"
)

type Svc struct{}

func Register(eg *echo.Group) {

	// 后台日志
	eg.Use(middlewares.LoggerMiddleware())
	// 不进行接口过滤
	no.Register(eg)

	// 接口过滤
	g := eg.Group("")
	//g.Use(middlewares.AuthMenu())
	//g.Use(middlewares.LoggerMiddleware())
	user.Register(g)
	menu.Register(g)
	adminlog.Register(g)

}
