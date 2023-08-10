package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/samber/lo"
	"github.com/spf13/viper"
)

const (
	envKey = "SERVICE_ENV"

	envProd = "prod"
	envTest = "test"
	envDev  = "dev"

	envDefault = envDev
)

const (
	configDefault = "config"

	configProd = envProd
	configTest = envTest
	configDev  = envDev
)

const (
	configBasePath = "internal/infrastructure/config/config"
)

var (
	envs = []string{envProd, envTest, envDev}
)

type Config struct {
	Service Service `mapstructure:"service"`
	Data    Data    `mapstructure:"data"`
	Log     Log     `mapstructure:"log"`
}

func Load() (Config, error) {
	env, err := getEnv()
	if err != nil {
		return lo.Empty[Config](), errors.Join(err, errors.New("load config"))
	}

	viper.AddConfigPath(configBasePath)

	c := Config{}

	err = readConfig(&c, configDefault)
	if err != nil {
		return lo.Empty[Config](), errors.Join(err, fmt.Errorf("read default config from: %s.yaml", configDefault))
	}

	configName, err := getConfigName(env)
	if err != nil {
		return lo.Empty[Config](), errors.Join(err, errors.New("convert env to config file name"))
	}

	err = readConfig(&c, configName)
	if err != nil {
		return lo.Empty[Config](), errors.Join(err, fmt.Errorf("read config from: %s.yaml", configName))
	}

	return c, nil
}

func getEnv() (string, error) {
	env, ok := os.LookupEnv(envKey)

	if !ok {
		return envDefault, nil
	}

	if !lo.Contains(envs, env) {
		return "", fmt.Errorf("environment variable %s should be set as one of %v", envKey, envs)
	}

	return env, nil
}

func getConfigName(env string) (string, error) {
	switch env {
	case envProd:
		return configProd, nil
	case envTest:
		return configTest, nil
	case envDev:
		return configDev, nil
	default:
		return lo.Empty[string](), fmt.Errorf("env is invalid: %s", env)
	}
}

func readConfig(c *Config, configName string) error {
	viper.SetConfigName(configName)

	err := viper.ReadInConfig()
	if err != nil {
		return errors.Join(err, fmt.Errorf("read in config from: %s.yaml", configName))
	}

	err = viper.Unmarshal(&c)
	if err != nil {
		return errors.Join(err, fmt.Errorf("unmarshal from: %s.yaml", configName))
	}

	return nil
}
