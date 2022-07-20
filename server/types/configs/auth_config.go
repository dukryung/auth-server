package configs

import (
	"encoding/json"
	"io/ioutil"
)

const (
	configPath = "./config.json"
)

type AuthConfig struct {
	Port int `json:"port"`
	DB DBConfig `json:"db"`
}

func LoadAuthConfig() (*AuthConfig, error) {

	var authConfig = &AuthConfig{}
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, authConfig)
	if err != nil {
		return nil, err
	}

	return authConfig, nil

}
