package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"fmt"
)

func Init(name string) {
	logLevel := map[string]log.Level{
		"info": log.InfoLevel,
		"warn": log.WarnLevel,
		"error": log.ErrorLevel,
		"fatal": log.FatalLevel,
	}

	viper.SetConfigName("config")
	viper.AddConfigPath(fmt.Sprintf("/etc/%s/", name))
	viper.AddConfigPath(fmt.Sprintf("$HOME/.%s/", name))
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config.yaml file: %s \n", err))
	}
	if level, ok := logLevel[viper.GetStringMapString("log")["level"]]; ok{
		log.SetLevel(level)
	} else {
		log.SetLevel(log.WarnLevel)
	}

}
