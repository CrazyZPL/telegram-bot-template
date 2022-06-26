package config

import (
	"io/ioutil"

	"golang.org/x/xerrors"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Bot   BotConfig
	Mysql MysqlConfig
}

type MysqlConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
}

type BotConfig struct {
	BotToken   string `yaml:"bot_token"`
	WebHookURL string `yaml:"webhook_url"`
}

func InitConfigFile(configFilePath string) (Config, error) {
	config := Config{}
	file, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return config, xerrors.Errorf("Fail to read config file: %v", err)
	}

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return config, xerrors.Errorf("Fail to unmarshal file: %v", err)
	}

	return config, nil
}
