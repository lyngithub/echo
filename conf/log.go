package conf

import "fmt"

type LogConfig struct {
	AdminPath   string `mapstructure:"admin_path"`
	AdminName   string `mapstructure:"admin_name"`
	AppPath     string `mapstructure:"app_path"`
	AppName     string `mapstructure:"app_name"`
	OpenapiPath string `mapstructure:"openapi_path"`
	OpenapiName string `mapstructure:"openapi_name"`
}

func (this *LogConfig) GetAdminFilePath() string {
	return fmt.Sprintf("%s/%s", Config.Log.AdminPath, Config.Log.AdminName)
}

func (this *LogConfig) GetAppFilePath() string {
	return fmt.Sprintf("%s/%s", Config.Log.AppPath, Config.Log.AppName)
}

func (this *LogConfig) GetOpenApiFilePath() string {
	return fmt.Sprintf("%s/%s", Config.Log.OpenapiPath, Config.Log.OpenapiName)
}
