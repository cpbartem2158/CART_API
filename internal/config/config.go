package config

import (
	"os"
	"time"

	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
	SSLMode  string `mapstructure:"sslmode"`
}

type ServerConfig struct {
	Port         string        `mapstructure:"port"`
	ReadTimeout  time.Duration `mapstructure:"read_timeout"`
	WriteTimeout time.Duration `mapstructure:"write_timeout"`
	IdleTimeout  time.Duration `mapstructure:"keepalive_timeout"`
}

type Config struct {
	Database DatabaseConfig `mapstructure:"database"`
	Server   ServerConfig   `mapstructure:"server"`
}

func LoadConfig(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	applyDatabaseEnvOverrides(&config)
	return &config, nil
}

func applyDatabaseEnvOverrides(cfg *Config) {
	if v := os.Getenv("POSTGRES_USER"); v != "" {
		cfg.Database.User = v
	}
	if v := os.Getenv("POSTGRES_PASSWORD"); v != "" {
		cfg.Database.Password = v
	}
	if v := os.Getenv("POSTGRES_DB"); v != "" {
		cfg.Database.DBName = v
	}
}
