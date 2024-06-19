package conf

import "fmt"

type XormConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DBname   string `mapstructure:"dbname"`
	ShowSql  bool   `mapstructure:"show_sql"`
}

func (this *XormConfig) GetDataSourceName() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		Config.Xorm.Username, Config.Xorm.Password, Config.Xorm.Host, Config.Xorm.Port, Config.Xorm.DBname)
}
