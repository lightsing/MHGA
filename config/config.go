package config

import (
	log "github.com/sirupsen/logrus"
	"fmt"
	"os"
	"path/filepath"
	"encoding/json"
	"io/ioutil"
	"errors"
)

const configName = "config.json"
var configPaths = []string {"/etc/mhga/", "$HOME/mhga/", "."}

var logLevel = map[string]log.Level{
	"info": log.InfoLevel,
	"warn": log.WarnLevel,
	"error": log.ErrorLevel,
	"fatal": log.FatalLevel,
}

type Config struct {
	Rules []struct {
		Name      string `json:"name"`
		Type      string `json:"type"`
		Upstream  string `json:"upstream"`
		Update    bool   `json:"update"`
		StorePath string `json:"store-path"`
		RulePath  string `json:"rule-path"`
	} `json:"rules"`
	Log struct {
		Level string `json:"level"`
	} `json:"log"`
}

func findConfig(name string) (*Config, error) {
	for _, path := range configPaths {
		path = filepath.Join(path, configName)
		fmt.Println(path)
		if _, err := os.Stat(path); err != nil {
			continue
		}
		if data, err := ioutil.ReadFile(path); err == nil {
			config := &Config{}
			if json.Unmarshal(data, config) == nil {
				return config, nil
			}
		}
	}
	return nil, errors.New("config not found")
}

func mustFindConfig(name string) *Config {
	config, err := findConfig(name)
	if err != nil {
		panic(err)
	}
	return config
}

func Init(name string) {
	config := mustFindConfig(name)
	switch config.Log.Level {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	case "panic":
		log.SetLevel(log.PanicLevel)
	case "quite":
		log.SetLevel(0)
	}

}
