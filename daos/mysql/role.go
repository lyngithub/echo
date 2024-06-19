package mysql

import (
	"echo/common/logger"
	"echo/daos"
	"echo/models"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"strconv"
)

type RoleDao struct {
}

func NewDaoRole() *RoleDao {
	return &RoleDao{}
}

func (m *RoleDao) UpdateBean(by int64, bean *models.Role) error {
	models.Edit(by, &bean.Bean)
	count, err := daos.Mysql.ID(bean.ID).Update(bean)
	if err != nil {
		return err
	}
	if 0 == count {
		return errors.New("Update 0 data")
	}
	return nil
}

func (m *RoleDao) InsertBean(by int64, bean *models.Role) error {
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

func (m *RoleDao) GetBeanById(id int64) (*models.Role, error) {
	bean := &models.Role{}
	has, err := daos.Mysql.ID(id).Get(bean)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("角色不存在")
	}
	return bean, nil
}

func (m *RoleDao) DeleteBeanById(id int64) error {
	_, err := daos.Mysql.ID(id).Delete(&models.Role{})
	return err
}

func (m *RoleDao) FindRoleByIds(ids []int64, status string) ([]*models.Role, error) {
	list := make([]*models.Role, 0)
	db := daos.Mysql.NewSession()
	if "" != status {
		db.Where("status = ?", status)
	}
	if nil != ids {
		db.In("id", ids)
	}
	err := db.Find(&list)
	return list, err
}

func (m *RoleDao) FindRoleIdsByUserId(userId int64) (list []int64) {
	sql := `
			SELECT
			r.id as id
		FROM
			a_user AS u,
			a_role AS r,
			a_user_role AS ur 
		WHERE
			u.status = "0" 
			AND r.status = "0" 
			AND ur.user_id = u.id 
			AND ur.role_id = r.id 
			and u.id = %d
`
	sql = fmt.Sprintf(sql, userId)
	v, err := daos.Mysql.QueryString(sql)
	if err != nil {
		logger.AdminLog.Error("[ADMIN] "+fmt.Sprintf("FindRoleIdsByUserId --> Fail [%d]", userId), zap.Error(err))
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

func (m *RoleDao) FindRoleIdsByRoleId(roleId int64) (list []int64) {
	sql := `
			SELECT
			r.id as id
		FROM
			a_user AS u,
			a_role AS r,
			a_user_role AS ur 
		WHERE
			u.status = "0" 
			AND r.status = "0" 
			AND ur.user_id = u.id 
			AND ur.role_id = r.id 
			and r.id = %d
`
	sql = fmt.Sprintf(sql, roleId)
	v, err := daos.Mysql.QueryString(sql)
	if err != nil {
		logger.AdminLog.Error("[ADMIN] "+fmt.Sprintf("FindRoleIdsByRoleId --> Fail [%d]", roleId), zap.Error(err))
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

func (m *RoleDao) DeleteUserRoleByUserId(userId int64) {
	daos.Mysql.Where("user_id = ?", userId).Delete(&models.UserRole{})
}

func (m *RoleDao) DeleteUserRoleByRoleId(userId int64) {
	daos.Mysql.Where("role_id = ?", userId).Delete(&models.UserRole{})
}

func (m *RoleDao) InsertUserRole(userRole *models.UserRole) error {
	count, err := daos.Mysql.Insert(userRole)
	if err != nil {
		return err
	}
	if 0 == count {
		return errors.New("Add 0 data")
	}
	return nil
}
