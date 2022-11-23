package config

import "github.com/spf13/viper"

type Config struct {
	ServerAddress    string `mapstructure:"SERVER_ADDRESS"`
	JwtKey           string `mapstructure:"JWT_KEY"`
	DBUri            string `mapstructure:"DB_URI"`
	AwsS3Bucket      string `mapstructure:"AWS_S3_BUCKET"`
	AwsS3Region      string `mapstructure:"AWS_S3_REGION"`
	AwsS3AccessKeyId string `mapstructure:"AWS_S3_ACCESS_KEY_ID"`
	AwsS3SecretKey   string `mapstructure:"AWS_S3_SECRET_KEY"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
