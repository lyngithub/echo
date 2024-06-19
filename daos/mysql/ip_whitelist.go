package mysql

import (
	"errors"
	"echo/daos"
	"echo/models"
)

type IpWhitelistDao struct {
}

func NewDaoIpWhitelist() *IpWhitelistDao {
	return &IpWhitelistDao{}
}

func (m *IpWhitelistDao) GetBeanById(id int64) (*models.IpWhitelist, error) {
	bean := &models.IpWhitelist{}
	has, err := daos.Mysql.ID(id).Get(bean)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("IpWhitelist does not exist")
	}
	return bean, nil
}

func (m *IpWhitelistDao) UpdateBean(by int64, bean *models.IpWhitelist) error {
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

func (m *IpWhitelistDao) DeleteBeanById(id, companyId int64) error {
	_, err := daos.Mysql.ID(id).Where("company_id = ?", companyId).Delete(&models.IpWhitelist{})
	return err
}

func (m *IpWhitelistDao) InsertBean(by int64, bean *models.IpWhitelist) error {
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

func (m *IpWhitelistDao) FindIpWhitelistAll() ([]*models.IpWhitelist, error) {
	list := make([]*models.IpWhitelist, 0)
	err := daos.Mysql.NewSession().Find(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (m *IpWhitelistDao) FindIpWhitelistById(companyId int64) ([]*models.IpWhitelist, error) {
	list := make([]*models.IpWhitelist, 0)
	err := daos.Mysql.Where("company_id = ?", companyId).Find(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}
