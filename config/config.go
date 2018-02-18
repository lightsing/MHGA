package config

import (
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path/filepath"
)

const configName = "config.json"

var configPaths = []string{"/etc/mhga/", "$HOME/mhga/", "."}

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
	Rules          []RuleConfig `json:"rules"`
	Address        string       `json:"address"`
	AvailableRules []string
	Log            struct {
		Level string `json:"level"`
	} `json:"log"`
}

func findConfig() (*Config, error) {
	for _, path := range configPaths {
		path = filepath.Join(path, configName)
		if _, err := os.Stat(path); err != nil {
			continue
		}
		if data, err := ioutil.ReadFile(path); err == nil {
			config := &Config{}
			if json.Unmarshal(data, config) == nil {
				log.Infof("found config at %s", path)
				return config, nil
			}
		}
	}
	return nil, errors.New("config not found")
}

func mustFindConfig() *Config {
	config, err := findConfig()
	if err != nil {
		panic(err)
	}
	return config
}

func Init() *Config {
	config := mustFindConfig()
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
	default:
		log.SetLevel(log.InfoLevel)
	}

	return config
}
