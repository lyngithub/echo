package services

import (
	"echo/common/logger"
	"echo/cons"
	"echo/daos/mysql"
	"echo/models"
	"echo/models/vo"
	"echo/utils"
	"fmt"
	"go.uber.org/zap"
)

type AdminlogService struct {
}

func NewServiceAdminlog() *AdminlogService {
	return &AdminlogService{}
}

func (s *AdminlogService) DeleteAdminlogByIds(ids []int64) error {
	for _, id := range ids {
		mysql.Adminlog.DeleteBeanById(id)
	}
	return nil
}

func (s *AdminlogService) FindAdminlog(userId int64, params *vo.FindListParams) (int64, []*vo.AdminlogVo, error) {
	u, err := mysql.User.GetBeanById(userId)
	if err != nil {
		logger.AdminLog.Error("[ADMIN] "+fmt.Sprintf("get user --> Fail [%d]", userId), zap.Error(err))
		return 0, nil, err
	}
	var (
		total        int64
		adminlogList []*models.Adminlog
	)
	if utils.IsAdmin(u) {
		total, adminlogList, err = mysql.Adminlog.FindAdminlogByIds(nil, params)
		if err != nil {
			logger.AdminLog.Error("[ADMIN] "+fmt.Sprintf("get a list of admin logs --> Fail [%v]", params), zap.Error(err))
			return 0, nil, err
		}
	} else {
		uids := mysql.User.FindUserIdsBySuperiorId(userId)
		if nil == uids {
			return 0, make([]*vo.AdminlogVo, 0), nil
		}
		total, adminlogList, err = mysql.Adminlog.FindAdminlogByIds(uids, params)
		if err != nil {
			logger.AdminLog.Error("[ADMIN] "+fmt.Sprintf("get a list of admin logs --> Fail [%d | %v]", uids, params), zap.Error(err))
			return 0, nil, err
		}
	}
	voList := make([]*vo.AdminlogVo, 0)
	for _, adminlog := range adminlogList {
		v := &vo.AdminlogVo{
			Id:        adminlog.ID,
			Method:    adminlog.Method,
			Url:       adminlog.Url,
			Params:    adminlog.Params,
			Name:      adminlog.Name,
			UserName:  adminlog.UserName,
			Ip:        adminlog.Ip,
			UserAgent: adminlog.UserAgent,
			CreateBy:  adminlog.CreateBy,
		}
		if !adminlog.Created.IsZero() {
			v.Created = adminlog.Created.Format(cons.TIMEDATETIME)
		}
		voList = append(voList, v)
	}
	return total, voList, nil
}
