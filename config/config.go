package config

import (
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path/filepath"
	"github.com/lightsing/makehttps/rules"
)

const configName = "config.json"

var configPaths = []string{"/etc/mhga/", "$HOME/mhga/", "."}

var logLevel = map[string]log.Level{
	"info":  log.InfoLevel,
	"warn":  log.WarnLevel,
	"error": log.ErrorLevel,
	"fatal": log.FatalLevel,
}

type GitConfig struct {
	Upstream     string `json:"upstream"`
	Path         string `json:"path"`
	Depth        int    `json:"depth"`
	SingleBranch bool   `json:"single-branch"`
	Branch       string `json:"branch"`
}

type RuleConfig struct {
	Name   string    `json:"name"`
	Type   string    `json:"type"`
	Update bool      `json:"update"`
	Git    GitConfig `json:"git"`
	Path   string    `json:"path"`
}

type Config struct {
	Rules []RuleConfig `json:"rules"`
	AvailableRules []string
	Log   struct {
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

func Init(name string) *Config {
	config := mustFindConfig(name)
	config.AvailableRules = make([]string, 0)

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

	for _, rule := range config.Rules {
		if err := rules.CheckRule(&rule); err == nil {
			config.AvailableRules = append(config.AvailableRules, rule.Path)
		} else {
			log.Errorf("Rule check fail by (%s)", err)
		}
	}

	return config
}
