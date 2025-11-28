package cmd

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	Token string `json:"token"`
}

func getConfigPath () string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".taskctl", "config.json")
}

func saveToken(token string) error {
	configPath := getConfigPath()
	os.Mkdir(filepath.Dir(configPath), 0755)
	config := Config{Token: token}
	configData, err := json.Marshal(config)
	if err != nil {
		return err
	}
	err = os.WriteFile(configPath, configData, 0644)
	return err
}

func loadToken() string {
	data, err := os.ReadFile(getConfigPath())
	if err != nil {
		return ""
	}
	var cfg Config
	json.Unmarshal(data, &cfg)
	return cfg.Token
}