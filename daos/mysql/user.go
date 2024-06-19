package mysql

import (
	"echo/common/logger"
	"echo/cons"
	"echo/daos"
	"echo/models"
	"echo/models/vo"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"strconv"
	"strings"
)

type UserDao struct {
}

func NewDaoUser() *UserDao {
	return &UserDao{}
}

func (m *UserDao) GetBeanById(id int64) (*models.User, error) {
	bean := &models.User{}
	has, err := daos.Mysql.ID(id).Where("status = ?", cons.STATUS_YES).Get(bean)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("Data does not exist")
	}
	return bean, nil
}

func (m *UserDao) GetUserById(id int64) (*models.User, error) {
	bean := &models.User{}
	has, err := daos.Mysql.ID(id).Get(bean)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("Data does not exist")
	}
	return bean, nil
}

func (m *UserDao) GetBeanByLoginName(loginName string) (*models.User, error) {
	bean := &models.User{}
	has, err := daos.Mysql.Where("login_name = ?", loginName).Get(bean)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return bean, nil
}

func (m *UserDao) UpdateBean(by int64, bean *models.User) error {
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

func (m *UserDao) FindAllUser() (list []*models.User, err error) {
	err = daos.Mysql.Where("status = ?", "0").Find(&list)
	return
}

func (m *UserDao) DeleteBeanById(id int64) error {
	_, err := daos.Mysql.ID(id).Delete(&models.User{})
	return err
}

func (m *UserDao) InsertBean(by int64, bean *models.User) error {
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

func (m *UserDao) FindUserByIds(ids []int64, param *vo.FindListParams) (int64, []*models.User, error) {
	list := make([]*models.User, 0)
	db := daos.Mysql.Limit(param.PageSize, (param.Page-1)*param.PageSize)
	if nil != ids {
		db.In("id", ids)
	}
	if "" != param.StartTime {
		db.Where("created >= ?", param.StartTime)
	}
	if "" != param.EndTime {
		db.Where("created <= ?", param.EndTime)
	}
	count, err := db.FindAndCount(&list)
	if err != nil {
		return 0, nil, err
	}
	return count, list, err
}

func (m *UserDao) FindUserIdsBySuperiorId(userId int64) (list []int64) {
	sql := `
			SELECT
				id 
			FROM
				a_user 
			WHERE
				id = ?
				OR superior_ids LIKE '%,?,%' 
				OR superior_ids LIKE '?,%' 
				OR superior_ids LIKE '%,?'
`
	sql = strings.ReplaceAll(sql, "?", fmt.Sprintf("%d", userId))
	v, err := daos.Mysql.QueryString(sql)
	if err != nil {
		logger.AdminLog.Error("[ADMIN] "+fmt.Sprintf("FindUserIdsBySuperiorId --> Fail [%d]", userId), zap.Error(err))
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
