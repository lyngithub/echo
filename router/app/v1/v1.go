package v1

import (
	"github.com/labstack/echo"
)

type Svc struct{}

func Register(eg *echo.Group) {
	_ = eg.Group("/v1")

}
