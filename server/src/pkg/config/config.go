package config

import (
	"encoding/json"
	"io/ioutil"
  l "github.com/FKouhai/OpenCaaServer/src/pkg/logger"
)

type Config struct {
	Etcd string `json:"Etcd"`
	Port string `json:"Port"`
}

func NewConfig() (*Config, error) {
	var config *Config
  newLog := l.NewLogger()
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
    l.LoggErr(newLog, err)
		return nil, err
	}
	err = json.Unmarshal(file, &config)
	if err != nil {
    l.LoggErr(newLog, err)
		return nil, err
	}
	return config, err
}
