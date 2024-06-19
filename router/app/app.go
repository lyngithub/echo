package app

import (
	"echo/middlewares"
	"github.com/labstack/echo"

	"echo/router/app/no"
	v1 "echo/router/app/v1"
)

type Svc struct{}

func Register(eg *echo.Group) {

	no.Register(eg)
	//auth.Register(eg)
	g := eg.Group("")
	g.Use(middlewares.AuthAppLogin())
	v1.Register(g)

}
