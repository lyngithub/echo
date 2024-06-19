package mysql

import (
	"echo/daos"
	"echo/models"
	"errors"
)

type VerDao struct {
}

func NewDaoVer() *VerDao {
	return &VerDao{}
}

func (m *VerDao) UpdateBean(bean *models.Ver) error {
	count, err := daos.Mysql.ID(bean.ID).Update(bean)
	if err != nil {
		return err
	}
	if 0 == count {
		return errors.New("Update 0 data")
	}
	return nil
}

func (m *VerDao) InsertBean(bean *models.Ver) error {
	count, err := daos.Mysql.Insert(bean)
	if err != nil {
		return err
	}
	if 0 == count {
		return errors.New("Add 0 data")
	}
	return nil
}

func (m *VerDao) GetBean() (*models.Ver, error) {
	bean := &models.Ver{}
	has, err := daos.Mysql.NewSession().Get(bean)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("not find")
	}
	return bean, nil
}
