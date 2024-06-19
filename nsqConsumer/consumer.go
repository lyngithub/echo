package nsqConsumer

import (
	"encoding/json"
	"fmt"
	"github.com/nsqio/go-nsq"
	"go.uber.org/zap"
	"echo/common/logger"
	"echo/conf"
	"echo/models/nsq_bean"
)

type MyTestHandler struct {
	q              *nsq.Consumer
	messageReceive int
}

func (h *MyTestHandler) HandleMessage(message *nsq.Message) error {
	defer message.Finish()
	data := new(nsq_bean.OrderDelay)
	err := json.Unmarshal(message.Body, data)
	if err != nil {
		logger.AdminLog.Error("[ADMIN] "+fmt.Sprintf("nsq获取推送信息失败 [msg:%v]", message), zap.Error(err))
	} else {
		if data.Topic == conf.Config.NsqProducer15MinConfig.TopicName {

		}
		if data.Topic == conf.Config.NsqProducer1HConfig.TopicName {

		}
	}
	return nil
}
func initConsuemr(addr, topic, channel string) {
	var (
		config *nsq.Config
		h      *MyTestHandler
		err    error
	)
	h = &MyTestHandler{}

	config = nsq.NewConfig()

	if h.q, err = nsq.NewConsumer(topic, channel, config); err != nil {
		logger.AdminLog.Error("[ADMIN] "+"注册nsq失败", zap.Error(err))
		panic(err)
	}

	h.q.AddHandler(h)
	if err = h.q.ConnectToNSQD(addr); err != nil {
		logger.AdminLog.Error("[ADMIN] "+"注册nsq失败", zap.Error(err))
		panic(err)
	}

	logger.AdminLog.Info("[ADMIN] " + "注册nsq成功，推送监听中...")
	return
}

func CreateConsuemr(addr, topic, channel string) {
	initConsuemr(addr, topic, channel)
	select {}
}
