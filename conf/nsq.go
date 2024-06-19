package conf

type NsqProducer15MinConfig struct {
	Has       bool   `mapstructure:"has"`
	NsqAddr   string `mapstructure:"nsq_addr"`
	TopicName string `mapstructure:"topic_name"`
}

type NsqConsumer15MinConfig struct {
	Has         bool   `mapstructure:"has"`
	NsqAddr     string `mapstructure:"nsq_addr"`
	TopicName   string `mapstructure:"topic_name"`
	ChannelName string `mapstructure:"channel_name"`
}

type NsqProducer1HConfig struct {
	Has       bool   `mapstructure:"has"`
	NsqAddr   string `mapstructure:"nsq_addr"`
	TopicName string `mapstructure:"topic_name"`
}

type NsqConsumer1HConfig struct {
	Has         bool   `mapstructure:"has"`
	NsqAddr     string `mapstructure:"nsq_addr"`
	TopicName   string `mapstructure:"topic_name"`
	ChannelName string `mapstructure:"channel_name"`
}
