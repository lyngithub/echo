package services

import (
	"echo/daos/mysql"
	"echo/models"
)

var (
	Menu     = NewServiceMenu()
	Role     = NewServiceRole()
	Adminlog = NewServiceAdminlog()
	User     = NewServiceUser()
)

var AdminLogChan chan *models.Adminlog

func init() {
	AdminLogChan = make(chan *models.Adminlog, 1000)
	go func() {
		for {
			select {
			case bean := <-AdminLogChan:
				// 查询菜单标题
				menu := mysql.Menu.GetBeanByMethodUrl(bean.Method, bean.Url)
				if nil != menu {
					if "S" == menu.Classification {
						continue
					}
					name := findParentMenuName(menu.ID)
					bean.Name = name
				}
				// 添加操作日志
				mysql.Adminlog.InsertBean(bean)
			}
		}
	}()
}

func findParentMenuName(mid int64) string {
	name := ""
	m, _ := mysql.Menu.GetBeanById(mid)
	if m != nil {
		name = m.MenuName
		if 0 != m.ParentId {
			name = findParentMenuName(m.ParentId) + " / " + name
		}
	}
	return name
}
