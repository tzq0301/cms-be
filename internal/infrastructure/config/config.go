package config

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/spf13/viper"
)

const (
	envKey = "APP_ENV"

	envProd = "prod"
	envTest = "test"
	envDev  = "dev"

	envDefault = envDev
)

var (
	envs = []string{envProd, envTest, envDev}
)

type Config struct {
	DB DB `mapstructure:"db"`
}

func Load() (Config, error) {
	env, err := getEnv()
	if err != nil {
		return lo.Empty[Config](), errors.Wrap(err, "cannot load config")
	}

	viper.SetConfigName(env)
	viper.AddConfigPath("internal/infrastructure/config")

	err = viper.ReadInConfig()
	if err != nil {
		return lo.Empty[Config](), errors.Wrap(err, "cannot read in config")
	}

	c := Config{}

	err = viper.Unmarshal(&c)
	if err != nil {
		return Config{}, errors.Wrapf(err, "cannot unmarshal from %s.yaml", env)
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
