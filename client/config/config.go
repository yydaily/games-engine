package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Server struct {
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	MaxDataBody int    `yaml:"max_data_body"`
}

type User struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Config struct {
	Server Server `yaml:"server"`
	User   User   `yaml:"user"`
}

var Conf *Config

func PraseConfig() {
	content, err := os.ReadFile("./config.yaml")
	if err != nil {
		panic(fmt.Sprintf("read config file failed, err=%s", err))
	}

	err = yaml.Unmarshal(content, &conf)
	if err != nil {
		panic(fmt.Errorf("parse config file failed, err=%s", err))
	}
}
