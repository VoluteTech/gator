package config

import (
	"encoding/json"
	"os"
)

const (
	configFileName = ".gatorconfig.json"
)

type Config struct {
	DBURL string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	fullPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	file, err := os.Open(fullPath)
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUserName = username
	return write(*cfg)
}

func getConfigFilePath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	configFilePath := homePath + "/" + configFileName
	return configFilePath, nil
}

func write(cfg Config) error {
	fullPath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}

	return nil
}
