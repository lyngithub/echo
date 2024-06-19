package driver

import (
	// user mysql
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

// MysqlConfig mysql配置参数
type MysqlConfig struct {
	DbDsn   string
	ShowSQL bool
}

// CreateMysql 初始化数据库组件
func CreateMysql(dbDsn string, showSql bool) (*xorm.Engine, error) {
	mysql, err := xorm.NewEngine("mysql", dbDsn)
	if err != nil {
		return nil, err
	}
	mysql.ShowSQL(showSql)
	err = mysql.Ping()
	if err != nil {
		panic(err)
	}
	return mysql, nil
}
