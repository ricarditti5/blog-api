package config

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
)

type Config struct {
	PORT         string
	DATABASE_URL string `mapstructure:"DATABASE_URL"`
}

func LoadConfig() (*Config, error) {
	err := gotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("\nError to load .env: %v ", err)
	}
	viper.SetDefault("PORT", "8080")
	viper.BindEnv("DATABASE_URL")
	viper.AutomaticEnv()

	var cfg Config

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("\nError to get variables ")
	}
	if cfg.DATABASE_URL == "" {
		return nil, fmt.Errorf("\nEmpty vabiables ")
	}
	return &cfg, nil
}
