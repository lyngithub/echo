package conf

type Wallet struct {
	Key string `mapstructure:"key"`
	Pk  string `mapstructure:"pk"`
	Url string `mapstructure:"url"`
}
