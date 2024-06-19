package common

import (
	"github.com/labstack/echo"
	"echo/conf"
	"net/http"
)

func (this *Svc) Ping(c echo.Context) error {
	c.Response().Header().Add("BuildTime", conf.BuildTime)
	c.Response().Header().Add("BuildHash", conf.BuildHash)
	return c.String(http.StatusOK, "pong")
}
