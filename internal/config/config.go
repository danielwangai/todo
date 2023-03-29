package config

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

const (
	envPrefix = "TODO_APP"
)

// WebServerConfig ...
type WebServerConfig struct {
	Port string `envconfig:"TODO_APP_SERVER_PORT" split_words:"true"`
}

type AppConfig struct {
	WebServer *WebServerConfig
}

// FromEnv loads the app config from environment variables
func FromEnv() (*AppConfig, error) {
	fromFileToEnv()
	cfg := &AppConfig{}
	if err := envconfig.Process(envPrefix, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func fromFileToEnv() { // determine config file loc, irrespective of the entry (main or test); it should resolve properly
	_, b, _, _ := runtime.Caller(0)
	cfgFilename := filepath.Join(filepath.Dir(b), "../../etc/config/config.local.env")
	fmt.Println("CFG: ", cfgFilename)

	if err := godotenv.Load(cfgFilename); err != nil {
		fmt.Printf("ERROR: Failure reading config file: %s\n", err)
	}
}
