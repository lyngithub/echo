package mysql

import (
	"errors"
	"echo/daos"
	"echo/models"
)

type LogininforDao struct {
}

func NewDaoLogininfor() *LogininforDao {
	return &LogininforDao{}
}

func (d *LogininforDao) Insert(bean *models.Logininfor) error {
	count, err := daos.Mysql.Insert(bean)
	if err != nil {
		return err
	}
	if 0 == count {
		return errors.New("Add 0 data")
	}
	return nil
}
