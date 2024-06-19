package services

import (
	"echo/common/logger"
	"echo/daos/mysql"
	"echo/utils/google"
	"fmt"
	"go.uber.org/zap"
)

type UserService struct {
}

func NewServiceUser() *UserService {
	return &UserService{}
}

func (s *UserService) ResetGoogleCode(by int64, uid int64) (string, string, error) {
	user, err := mysql.User.GetUserById(uid)
	if err != nil {
		logger.AdminLog.Error("[ADMIN] "+fmt.Sprintf("get user --> Fail [%d]", uid), zap.Error(err))
		return "", "", err
	}
	googleCode, err := google.CreateGoogleSecret()
	if err != nil {
		logger.AdminLog.Error("[ADMIN] "+fmt.Sprintf("create company google_code --> Fail [%d]", uid), zap.Error(err))
		return "", "", err
	}
	user.GoogleCode = googleCode

	err = mysql.User.UpdateBean(by, user)
	if err != nil {
		logger.AdminLog.Error("[ADMIN] "+fmt.Sprintf("edit user google_code --> Fail [%v]", user), zap.Error(err))
		return "", "", err
	}
	return user.Username, googleCode, nil
}
