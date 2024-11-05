package helper

import (
	"github.com/spf13/viper"
)

type Config struct {
	DB          string `mapstructure:"DB"`
	SCHEMA      string `mapstructure:"SCHEMA"`
	GinMode     string `mapstructure:"GIN_MODE"`
	Env         string `mapstructure:"ENV"`
	LogFile     string `mapstructure:"LOG_FILE"`
	AutoMigrate string `mapstructure:"AUTO_MIGRATE"`
	Port        string `mapstructure:"PORT"`
	AllowOrigin string `mapstructure:"ALLOW_ORIGIN"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigName("app")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
