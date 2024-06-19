package mysql

import (
	"echo/daos"
	"echo/models"
	"echo/models/vo"
	"errors"
)

type AdminlogDao struct {
}

func NewDaoAdminlog() *AdminlogDao {
	return &AdminlogDao{}
}

func (m *AdminlogDao) InsertBean(bean *models.Adminlog) error {
	count, err := daos.Mysql.Insert(bean)
	if err != nil {
		return err
	}
	if 0 == count {
		return errors.New("Add 0 data")
	}
	return nil
}

func (m *AdminlogDao) DeleteBeanById(id int64) error {
	_, err := daos.Mysql.ID(id).Delete(&models.Adminlog{})
	return err
}

func (m *AdminlogDao) FindAdminlogByIds(ids []int64, param *vo.FindListParams) (int64, []*models.Adminlog, error) {
	list := make([]*models.Adminlog, 0)
	db := daos.Mysql.Limit(param.PageSize, (param.Page-1)*param.PageSize).Desc("id")
	if nil != ids {
		db.In("create_by", ids)
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
