package conf

type LeveldbConfig struct {
	Has  bool   `mapstructure:"has"`
	Path string `mapstructure:"path"`
}
