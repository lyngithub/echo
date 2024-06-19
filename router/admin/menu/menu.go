package menu

import (
	"echo/common/logger"
	"echo/models/resp"
	"echo/services"
	"fmt"
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"strconv"
	"strings"
)

type Svc struct{}

func Register(eg *echo.Group) {
	s := &Svc{}
	g := eg.Group("/menu")
	{

		g.POST("/del", s.Del)
	}
}

func (s *Svc) Del(c echo.Context) error {
	var params struct {
		Ids string `json:"ids"`
	}
	if err := c.Bind(&params); err != nil {
		return resp.Fail(c, "Parameter error")
	}
	if "" == params.Ids {
		return resp.Fail(c, "Parameter error")
	}
	menuIds := make([]int64, 0)
	split := strings.Split(params.Ids, ",")
	for _, id := range split {
		userId, err := strconv.Atoi(id)
		if err != nil {
			logger.AdminLog.Error("[ADMIN] "+fmt.Sprintf("delete menu --> Fail [%s]", params.Ids), zap.Error(err))
			return resp.Fail(c, "Parameter error")
		}
		menuIds = append(menuIds, int64(userId))
	}
	err := services.Menu.DeleteMenuByIds(menuIds)
	if err != nil {
		logger.AdminLog.Error("[ADMIN] "+fmt.Sprintf("delete menu --> Fail [%v]", params), zap.Error(err))
		return resp.Fail(c, err.Error())
	}
	return resp.OK(c, "")
}
