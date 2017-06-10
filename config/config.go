package config

import (
	"fmt"
	env "github.com/joho/godotenv"
	"os"
	"strings"
)

func InitConfig(configPath string) error {
	if strings.TrimSpace(configPath) == "" {
		return fmt.Errorf("%s: %s", "Error", "cannot find config path")
		os.Exit(1)
	}
	configPath = strings.TrimRight(configPath, "/")

	err := env.Load()
	if err != nil {
		return fmt.Errorf("%s: %s", "Error", err.Error())
	}

	os.Setenv("DIR", configPath)
	return nil
}
