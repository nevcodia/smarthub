package bootstrap

import (
	"github.com/spf13/viper"
	"log"
)

type Env struct {
	AppEnv      string `mapstructure:"APP_ENV"`
	Host        string `mapstructure:"HOST"`
	Port        string `mapstructure:"PORT"`
	S3HostAddr  string `mapstructure:"S3_HOST_ADDR"`
	S3AccessKey string `mapstructure:"S3_ACCESS_KEY"`
	S3SecretKey string `mapstructure:"S3_SECRET_KEY"`
	S3Region    string `mapstructure:"S3_REGION"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if env.AppEnv == "development" {
		log.Println("The App is running in development env")
	}

	return &env
}
