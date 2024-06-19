package conf

type JwtConfig struct {
	Expire int    `mapstructure:"expire"`
	Secret string `mapstructure:"secret"`
}
