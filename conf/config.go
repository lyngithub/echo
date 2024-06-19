package conf

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

var (
	Config      = new(Configuration)
	BuildTime   = ""
	BuildHash   = ""
	MinAdAmount = make(map[int64]string)
)

type Configuration struct {
	Echo                   EchoConfig             `mapstructure:"echo"`
	Jwt                    JwtConfig              `mapstructure:"jwt"`
	Xorm                   XormConfig             `mapstructure:"mysql"`
	Redis                  RedisConfig            `mapstructure:"redis"`
	Leveldb                LeveldbConfig          `mapstructure:"leveldb"`
	NsqProducer15MinConfig NsqProducer15MinConfig `mapstructure:"nsq_producer_15min"`
	NsqConsumer15MinConfig NsqConsumer15MinConfig `mapstructure:"nsq_consumer_15min"`
	NsqProducer1HConfig    NsqProducer1HConfig    `mapstructure:"nsq_producer_1h"`
	NsqConsumer1HConfig    NsqConsumer1HConfig    `mapstructure:"nsq_consumer_1h"`
	Log                    LogConfig              `mapstructure:"log"`
	AliyunOSS              AliyunOSS              `mapstructure:"alioss"`
	Wallet                 Wallet                 `mapstructure:"wallet"`
	Chat                   Chat                   `mapstructure:"chat"`
	CoinPrice              CoinPrice              `mapstructure:"coin_price"`
	Airdrop                Airdrop                `mapstructure:"airdrop"`
	Cashier                Cashier                `mapstructure:"cashier"`
}

func InitializeConfig(path string) *viper.Viper {
	if configEnv := os.Getenv("VIPER_CONFIG"); configEnv != "" {
		path = configEnv
	}

	v := viper.New()
	v.SetConfigFile(path)
	v.SetConfigType("toml")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read config failed: %s", err))
	}
	v.WatchConfig()

	readConfig := func() {
		if err := v.Unmarshal(Config); err != nil {
			fmt.Println(err)
		}
	}

	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
		readConfig()
	})
	readConfig()
	return v
}
