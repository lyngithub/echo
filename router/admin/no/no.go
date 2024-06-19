package no

import (
	"echo/common/logger"
	"echo/models/resp"
	"echo/services"
	"fmt"
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"strconv"
)

type Svc struct{}

func Register(eg *echo.Group) {
	s := &Svc{}
	g := eg.Group("/no")
	{
		g.GET("/menu", s.Menu)                         // 获取菜单树
		g.GET("/role_menu/:rid", s.RoleMenu)           // 根据角色获取菜单
		g.GET("/roletree", s.RoleTree)                 // 获取角色树
		g.GET("/role", s.Role)                         // 获取角色列表
		g.GET("/find_menu/:mid", s.FindMenuByParentId) // 根据菜单id获取子菜单列表
	}
}

func (s *Svc) Menu(c echo.Context) error {
	//userId := utils.UserId(c)
	userId := int64(1)
	list, err := services.Menu.GetMenuTree(userId)
	if err != nil {
		logger.AdminLog.Error("[ADMIN] "+fmt.Sprintf("get menu tree --> Fail [%d]", userId), zap.Error(err))
		return resp.Fail(c, err.Error())
	}
	m := make(map[string]interface{})
	m["list"] = list
	return resp.OK(c, m)
}

func (s *Svc) RoleMenu(c echo.Context) error {
	param := c.Param("rid")
	rid, err := strconv.Atoi(param)
	if err != nil {
		return resp.Fail(c, "Parameter error")
	}
	list, err := services.Menu.GetRoleMenuTree(int64(rid))
	if err != nil {
		logger.AdminLog.Error("[ADMIN] "+fmt.Sprintf("get character menu tree --> Fail [%d]", rid), zap.Error(err))
		return resp.Fail(c, "fetch failed")
	}
	m := make(map[string]interface{})
	m["list"] = list
	return resp.OK(c, m)
}

func (s *Svc) RoleTree(c echo.Context) error {
	//userId := utils.UserId(c)
	userId := int64(1)
	list, err := services.Role.GetRoleTree(userId, "0")
	if err != nil {
		logger.AdminLog.Error("[ADMIN] "+fmt.Sprintf("get character tree --> Fail [%d]", userId), zap.Error(err))
		return resp.Fail(c, "fetch failed")
	}
	m := make(map[string]interface{})
	m["list"] = list
	return resp.OK(c, m)
}

func (s *Svc) Role(c echo.Context) error {
	//userId := utils.UserId(c)
	userId := int64(1)
	list, err := services.Role.FindRole(userId, "0")
	if err != nil {
		logger.AdminLog.Error("[ADMIN] "+fmt.Sprintf("get role --> Fail [%d]", userId), zap.Error(err))
		return resp.Fail(c, "fetch failed")
	}
	m := make(map[string]interface{})
	m["list"] = list
	return resp.OK(c, m)
}

func (s *Svc) FindMenuByParentId(c echo.Context) error {
	param := c.Param("mid")
	mid, err := strconv.Atoi(param)
	if err != nil {
		return resp.Fail(c, "Parameter error")
	}
	//userId := utils.UserId(c)
	userId := int64(1)
	list, err := services.Menu.FindMenuByParentId(userId, int64(mid))
	if err != nil {
		logger.AdminLog.Error("[ADMIN] "+fmt.Sprintf("get submenu --> Fail [%d]", mid), zap.Error(err))
		return resp.Fail(c, "fetch failed")
	}
	m := make(map[string]interface{})
	m["list"] = list
	return resp.OK(c, m)
}
