package driver

import (
	"github.com/nsqio/go-nsq"
	"go.uber.org/zap"
	"echo/common/logger"
)

type NsqProducerConfig struct {
	Addr  string
	Topic string
}

func CreateProduer(addr, topic string) (*nsq.Producer, string, error) {
	c := nsq.NewConfig()
	//err := c.Set("max_requeue_delay", "1440m")
	//err := c.Set("default_requeue_delay", "1440m")
	//if err != nil || err != nil {
	//	logger.AdminLog.Error("[ADMIN] "+"修改max_requeue_delay 失败", zap.Error(err))
	//}
	p, err := nsq.NewProducer(addr, c)
	if err != nil {
		logger.AdminLog.Error("[ADMIN] "+"启动注册nsq失败", zap.Error(err))
		panic(err)
	}
	return p, topic, nil
}
