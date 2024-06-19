package conf

type AliyunOSS struct {
	AccessKeyId     string `mapstructure:"access_key_id"`
	AccessKeySecret string `mapstructure:"access_key_secret"`
	Endpoint        string `mapstructure:"endpoint"`
	BucketName      string `mapstructure:"bucket_name"`
	BasePath        string `mapstructure:"base_path"`
	PhotoHost       string `mapstructure:"photo_host"`
}
