package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

func InitConfig() {
	wd, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(wd + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("init config failed")
	}
}
