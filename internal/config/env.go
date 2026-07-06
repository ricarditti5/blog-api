package config

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
)

type Config struct {
	PORT   string
	DB_URL string
}

func LoadConfig() (*Config, error) {
	err := gotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("Error to load .env")
	}
	viper.SetDefault("PORT", "8080")
	viper.BindEnv("DATABASE_URL")
	viper.AutomaticEnv()

	var cfg Config

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("Error to get variables")
	}
	if cfg.DB_URL == "" {
		return nil, fmt.Errorf("Empty vabiables")
	}
	return &cfg, nil
}
