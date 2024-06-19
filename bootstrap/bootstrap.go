package bootstrap

import (
	"echo/common/driver"
	"echo/common/logger"
	"echo/conf"
	"echo/daos"
	"echo/job"
	"echo/nsqConsumer"
	"echo/nsqProducer"
	"echo/systemSetting"
	"echo/utils/jwt_auth"
	"github.com/go-xorm/xorm"

	"time"
)

func Bootstrap() {
	systemSetting.CheckInit()
	// 设置log
	logger.InitAdminLog(conf.Config.Log.GetAdminFilePath())
	logger.InitAppLog(conf.Config.Log.GetAppFilePath())
	logger.InitOpenApiLog(conf.Config.Log.GetOpenApiFilePath())
	initMysql()
	go job.InitConsole()
	if conf.Config.Redis.Has {
		initRedis()
	}
	if conf.Config.Leveldb.Has {
		initLevelDB()
	}
	if conf.Config.NsqProducer15MinConfig.Has {
		initNsq15MinProducer()
	}
	if conf.Config.NsqConsumer15MinConfig.Has {
		go func() {
			initNsq15MinConsumer()
		}()
	}

	if conf.Config.NsqProducer1HConfig.Has {
		initNsq1HProducer()
	}
	if conf.Config.NsqConsumer1HConfig.Has {
		go func() {
			initNsq1HConsumer()
		}()
	}
	// jwt
	jwt_auth.EXPIRE = time.Duration(conf.Config.Jwt.Expire) * time.Hour
	jwt_auth.SECRET = conf.Config.Jwt.Secret
	// 版本检查
	err := InitVer()
	if err != nil {
		panic(err)
	}

}

func initMysql() {
	// 获取mysql配置
	engine, err := driver.CreateMysql(conf.Config.Xorm.GetDataSourceName(), conf.Config.Xorm.ShowSql)
	if err != nil {
		panic(err)
	}
	engine.SetMaxIdleConns(5000)
	engine.SetMaxOpenConns(5000)
	engine.SetLogger(xorm.NewSimpleLogger(logger.Xlogger))
	engine.ShowSQL(conf.Config.Xorm.ShowSql)
	daos.SetMysql(engine)
}

func initRedis() {
	// 获取redis配置
	rc, err := driver.CreateRedis(conf.Config.Redis.GetAddr(), conf.Config.Redis.Password, conf.Config.Redis.DB)
	if err != nil {
		panic(err)
	}
	daos.SetRedis(rc)
}

func initLevelDB() {
	leveldb, err := driver.CreateLevelDB(conf.Config.Leveldb.Path)
	if err != nil {
		panic(err)
	}
	daos.SetLevelDB(leveldb)
}

func initNsq15MinProducer() {
	var err error
	nsqProducer.Producer15Min, _, err = driver.CreateProduer(conf.Config.NsqProducer15MinConfig.NsqAddr, conf.Config.NsqProducer15MinConfig.TopicName)
	if err != nil {
		panic(err)
	}
	err = nsqProducer.Producer15Min.Ping()
	if err != nil {
		panic(err)
	}
}

func initNsq1HProducer() {
	var err error
	nsqProducer.Producer1H, _, err = driver.CreateProduer(conf.Config.NsqProducer1HConfig.NsqAddr, conf.Config.NsqProducer1HConfig.TopicName)
	if err != nil {
		panic(err)
	}
	err = nsqProducer.Producer1H.Ping()
	if err != nil {
		panic(err)
	}
}

func initNsq15MinConsumer() {
	nsqConsumer.CreateConsuemr(conf.Config.NsqConsumer15MinConfig.NsqAddr, conf.Config.NsqConsumer15MinConfig.TopicName, conf.Config.NsqConsumer15MinConfig.ChannelName)
}

func initNsq1HConsumer() {
	nsqConsumer.CreateConsuemr(conf.Config.NsqConsumer1HConfig.NsqAddr, conf.Config.NsqConsumer1HConfig.TopicName, conf.Config.NsqConsumer1HConfig.ChannelName)
}
