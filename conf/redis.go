package conf

import "fmt"

type RedisConfig struct {
	Has      bool   `mapstructure:"has"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

func (this *RedisConfig) GetAddr() string {
	return fmt.Sprintf("%s:%s",
		Config.Redis.Host, Config.Redis.Port)
}
