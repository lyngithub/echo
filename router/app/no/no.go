package no

import (
	"github.com/labstack/echo"
)

type Svc struct{}

func Register(eg *echo.Group) {
	s := &Svc{}
	g := eg.Group("/no")
	{
		g.GET("/test", s.test)
	}
}

func (s *Svc) test(echo.Context) error {
	return nil
}
