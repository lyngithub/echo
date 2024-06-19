package daos

import (
	"github.com/go-xorm/xorm"
	"echo/common/driver"
	"gopkg.in/redis.v5"
	"log"
	"time"
)

var (
	Mysql *xorm.Engine
	Rc    *redis.Client
	Ldb   *driver.LevelDB
)

// SetLevelDB
func SetLevelDB(_leveldb *driver.LevelDB) {
	Ldb = _leveldb
}

// SetMysql
func SetMysql(e *xorm.Engine) {
	Mysql = e
	go func() {
		for {
			Mysql.Ping()
			time.Sleep(1 * time.Hour)
		}
	}()
}

// SetRedis
func SetRedis(_rc *redis.Client) {
	Rc = _rc
}

// Transaction
func Transaction(fs ...func(s *xorm.Session) error) error {
	session := Mysql.NewSession()
	session.Begin()
	for _, f := range fs {
		err := f(session)
		if err != nil {
			log.Println(err)
			session.Rollback()
			session.Clone()
			return err
		}
	}
	session.Commit()
	session.Clone()
	return nil
}

// CloseMySQL
func CloseMySQL() {
	Mysql.Close()
}
