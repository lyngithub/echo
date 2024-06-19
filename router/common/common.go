package common

import (
	"github.com/labstack/echo"
	"net/http"
)

type Svc struct{}

func Register(e *echo.Group) {
	s := &Svc{}
	e.GET("/", s.Index)
	e.HEAD("/ping", s.Ping)
}

func (this *Svc) Index(c echo.Context) error {
	return c.String(http.StatusOK, "hello world")
}
