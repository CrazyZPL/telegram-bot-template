package utils

import (
	"io/ioutil"

	"golang.org/x/xerrors"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Bot BotConfig
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
