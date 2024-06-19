package adminlog

import (
	"echo/models/resp"
	"github.com/labstack/echo"
)

type Svc struct{}

func Register(eg *echo.Group) {
	s := &Svc{}
	g := eg.Group("/adminlog")
	{
		g.POST("/find", s.Find)

	}
}

func (s *Svc) Find(c echo.Context) error {

	m := make(map[string]interface{})

	return resp.OK(c, m)
}
