package mysql

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"echo/common/logger"
	"echo/daos"
	"echo/models"
	"strconv"
)

type MenuDao struct {
}

func NewDaoMenu() *MenuDao {
	return &MenuDao{}
}

func (m *MenuDao) UpdateBean(by int64, bean *models.Menu) error {
	models.Edit(by, &bean.Bean)
	count, err := daos.Mysql.ID(bean.ID).AllCols().Update(bean)
	if err != nil {
		return err
	}
	if 0 == count {
		return errors.New("Update 0 data")
	}
	return nil
}

func (m *MenuDao) GetBeanById111(id int64) (bool, *models.Menu, error) {
	bean := &models.Menu{}
	has, err := daos.Mysql.ID(id).Get(bean)
	if err != nil {
		return false, nil, err
	}

	return has, bean, nil
}

func (m *MenuDao) GetBeanById(id int64) (*models.Menu, error) {
	bean := &models.Menu{}
	has, err := daos.Mysql.ID(id).Get(bean)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("菜单不存在")
	}
	return bean, nil
}

func (m *MenuDao) GetBeanByMenuName11(menuName string, parentId int64) (bool, *models.Menu, error) {
	bean := &models.Menu{}
	has, err := daos.Mysql.Where("menu_name = ?", menuName).And("parent_id = ?", parentId).Get(bean)
	if err != nil {
		return false, nil, err
	}

	return has, bean, nil
}

func (m *MenuDao) GetBeanByMenuName(menuName string, parentId int64) (*models.Menu, error) {
	bean := &models.Menu{}
	has, err := daos.Mysql.Where("menu_name = ?", menuName).And("parent_id = ?", parentId).Get(bean)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("菜单不存在")
	}
	return bean, nil
}

func (m *MenuDao) InsertBean(by int64, bean *models.Menu) error {
	models.Add(by, &bean.Bean)
	count, err := daos.Mysql.Insert(bean)
	if err != nil {
		return err
	}
	if 0 == count {
		return errors.New("Add 0 data")
	}
	return nil
}

func (m *MenuDao) DeleteBeanById(id int64) error {
	_, err := daos.Mysql.ID(id).Delete(&models.Menu{})
	return err
}

func (m *MenuDao) InsertRoleMenu(userRole *models.RoleMenu) error {
	count, err := daos.Mysql.Insert(userRole)
	if err != nil {
		return err
	}
	if 0 == count {
		return errors.New("Add 0 data")
	}
	return nil
}

func (m *MenuDao) DeleteRoleMenuByRoleId(roleId int64) {
	daos.Mysql.Where("role_id = ?", roleId).Delete(&models.RoleMenu{})
}

func (m *MenuDao) DeleteRoleMenuByMenuId(menuId int64) {
	daos.Mysql.Where("menu_id = ?", menuId).Delete(&models.RoleMenu{})
}

func (m *MenuDao) GetBeanByMethodUrl(method, url string) *models.Menu {
	bean := &models.Menu{}
	has, err := daos.Mysql.Where("method = ?", method).And("url = ?", url).Get(bean)
	if err != nil || !has {
		return nil
	}
	return bean
}

func (m *MenuDao) FindMenuByIds(ids []int64, visible string, parentId int64) ([]*models.Menu, error) {
	list := make([]*models.Menu, 0)
	db := daos.Mysql.NewSession().Asc("id")
	if "" != visible {
		db.Where("visible = ?", visible)
	}
	if 0 != parentId {
		db.Where("parent_id = ?", parentId)
	}
	if nil != ids {
		db.In("id", ids)
	}
	err := db.Find(&list)
	return list, err
}

func (m *MenuDao) FindMenuIdsByUserId(userId int64) (list []int64) {
	sql := `
			SELECT
			m.id as id
		FROM
			a_user AS u,
			a_role AS r,
			a_menu AS m,
			a_role_menu AS rm,
			a_user_role AS ur 
		WHERE
			u.status = "0" 
			AND r.status = "0" 
			AND ur.user_id = u.id 
			AND ur.role_id = r.id 
			AND rm.role_id = r.id 
			AND rm.menu_id = m.id
			and u.id = %d
`
	sql = fmt.Sprintf(sql, userId)
	v, err := daos.Mysql.QueryString(sql)
	if err != nil {
		logger.AdminLog.Error("[ADMIN] "+fmt.Sprintf("FindMenuIdsByUserId --> Fail [%d]", userId), zap.Error(err))
		return
	}
	list = make([]int64, 0)
	for _, m := range v {
		id := m["id"]
		atoi, err := strconv.Atoi(id)
		if err != nil {
			continue
		}
		list = append(list, int64(atoi))
	}
	return
}

func (m *MenuDao) FindMenuIdsByRoleId(roleId int64) (list []int64) {
	sql := `
			SELECT
			m.id as id
		FROM
			a_role AS r,
			a_menu AS m,
			a_role_menu AS rm
		WHERE
			r.status = "0" 
			AND rm.role_id = r.id 
			AND rm.menu_id = m.id
			and r.id = %d
`
	sql = fmt.Sprintf(sql, roleId)
	v, err := daos.Mysql.QueryString(sql)
	if err != nil {
		logger.AdminLog.Error("[ADMIN] "+fmt.Sprintf("FindMenuIdsByRoleId --> Fail [%d]", roleId), zap.Error(err))
		return
	}
	list = make([]int64, 0)
	for _, m := range v {
		id := m["id"]
		atoi, err := strconv.Atoi(id)
		if err != nil {
			continue
		}
		list = append(list, int64(atoi))
	}
	return
}

func (m *MenuDao) FindMenuIdsByMenuId(menuId int64) (list []int64) {
	sql := `
			SELECT
			m.id as id
		FROM
			a_role AS r,
			a_menu AS m,
			a_role_menu AS rm
		WHERE
			 rm.role_id = r.id 
			AND rm.menu_id = m.id
			and m.id = %d
			and r.id <> 1
`
	sql = fmt.Sprintf(sql, menuId)
	v, err := daos.Mysql.QueryString(sql)
	if err != nil {
		logger.AdminLog.Error("[ADMIN] "+fmt.Sprintf("FindMenuIdsByMenuId --> Fail [%d]", menuId), zap.Error(err))
		return
	}
	list = make([]int64, 0)
	for _, m := range v {
		id := m["id"]
		atoi, err := strconv.Atoi(id)
		if err != nil {
			continue
		}
		list = append(list, int64(atoi))
	}
	return
}
