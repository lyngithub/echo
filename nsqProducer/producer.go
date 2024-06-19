package nsqProducer

import (
	"echo/common/logger"
	"github.com/nsqio/go-nsq"
	"go.uber.org/zap"
	"time"
)

var Producer15Min *nsq.Producer
var Producer1H *nsq.Producer

type IProducer interface {
	Publish(message string) error
	DeferredPublish(delay time.Duration, message string) error
}
type ProducerManager struct {
	producer *nsq.Producer
	topic    string
}

func Producer(p *nsq.Producer, topic string) IProducer {
	return &ProducerManager{producer: p, topic: topic}
}

func (p *ProducerManager) Publish(message string) (err error) {
	if message == "" {
		logger.AdminLog.Error("[ADMIN] " + "nsq推送消息为nil")
		return nil
	}
	if err = p.producer.Publish(p.topic, []byte(message)); err != nil {
		logger.AdminLog.Error("[APP] "+"nsq推送消息 失败", zap.Error(err))
		return err
	}
	return nil
}

// 延迟消息
func (p *ProducerManager) DeferredPublish(delay time.Duration, message string) (err error) {
	if message == "" {
		logger.AdminLog.Error("[APP] " + "nsq推送延迟消息为nil")
		return nil
	}
	if err = p.producer.DeferredPublish(p.topic, delay, []byte(message)); err != nil {
		logger.AdminLog.Error("[ADMIN] "+"nsq推送延迟消息 失败", zap.Error(err))
		return err
	}
	return nil
}

//func DelayPublish(delay time.Duration, param *nsq_bean.OrderDelay) error {
//	var p *nsq.Producer
//	if delay == cons.OrderDelay15Min {
//		p = Producer15Min
//	}
//	if delay == cons.OrderDelay1H {
//		p = Producer1H
//	}
//	if err := p.Ping(); err != nil {
//		return err
//	}
//	msg, err := json.Marshal(param)
//	if err != nil {
//		return err
//	}
//	if err := Producer(p, param.Topic).DeferredPublish(delay, string(msg)); err != nil {
//		return err
//	}
//	return nil
//}
