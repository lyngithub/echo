package bootstrap

import (
	"echo/common/logger"
	"echo/daos"
	"echo/daos/mysql"
	"echo/models"
	"errors"
	"fmt"
	"go.uber.org/zap"
)

func InitVer() error {
	newVer := "v1.76"
	// 所有版本
	vers := []string{"v1.76"}
	model := &models.Ver{}
	has, _ := daos.Mysql.IsTableExist(model)
	if !has {
		_ = daos.Mysql.CreateTables(model)
	}
	// 获取版本数据
	model, err := mysql.Ver.GetBean()
	if nil == model {
		model = &models.Ver{}
		model.Key = "v1.0"
		if err = mysql.Ver.InsertBean(model); err != nil {
			logger.AdminLog.Error("[ADMIN] "+fmt.Sprintf("初始化版本 错误 [%v]", model), zap.Error(err))
			return errors.New("初始化版本 错误")
		}
	}
	if "" == model.Key {
		model.Key = "v1.0"
		if err = mysql.Ver.UpdateBean(model); err != nil {
			logger.AdminLog.Error("[ADMIN] "+fmt.Sprintf("修复版本 错误 [%v]", model), zap.Error(err))
			return errors.New("修复版本 错误")
		}
	}

	for _, ver := range vers {
		if newVer == model.Key {
			break
		}
		switch ver {
		case "v1.76":
			v96(model, ver)
		}
	}
	return nil
}

func v96(model *models.Ver, verKey string) {

	sql := `SELECT TRIGGER_NAME FROM information_schema.TRIGGERS WHERE  TRIGGER_NAME = ?`
	var triggerName string
	has, err := daos.Mysql.SQL(sql, "update_floating_price1").Get(&triggerName)
	if has {
		return
	}
	// 创建触发器
	sqlTrigger := `
		CREATE TRIGGER update_floating_price1
		AFTER UPDATE ON a_coin
		FOR EACH ROW
		BEGIN
		    IF NEW.floating_buy != OLD.floating_buy OR NEW.floating_sell != OLD.floating_sell THEN
		        UPDATE a_merchant_advertisement AS ma
		        JOIN a_coin AS ac ON ma.coin_id = ac.id
		        SET ma.price = CASE
		                                    WHEN ma.type = 1 THEN ma.floating_ratio * ac.floating_sell
		                                    WHEN ma.type = 2 THEN ma.floating_ratio * ac.floating_buy
		                                    ELSE ma.price
		                                END,
		            ma.min_number = IF(ma.price <> 0, ma.min_limit / ma.price, ma.min_number),
		            ma.max_number = IF(ma.price <> 0, ma.max_limit / ma.price, ma.max_number),
		            ma.price_order = IF(ma.price <> 0, ma.price, ma.price_order)
		        WHERE ma.style = 2;
		    END IF;
		END;
	`
	_, err = daos.Mysql.Exec(sqlTrigger)
	if err != nil {
		return
	}
	if err == nil {
		// 修改数据库
		model.Key = verKey
		mysql.Ver.UpdateBean(model)
	}
}
