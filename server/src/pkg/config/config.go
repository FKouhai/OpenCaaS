package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Etcd string `json:"Etcd"`
	Port string `json:"Port"`
}

func NewConfig() (*Config, error) {
	var config *Config
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Printf("Error reading config %s", err)
		return nil, err
	}
	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println("Error unmarshalling config", err.Error())
		return nil, err
	}
	return config, err
}
