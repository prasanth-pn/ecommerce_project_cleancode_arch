package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	DBHost           string `mapstructure:"DB_HOST"`
	DBName           string `mapstructure:"DB_NAME"`
	DBUser           string `mapstructure:"DB_USER"`
	DBPort           string `mapstructure:"DB_PORT"`
	DBPassword       string `mapstructure:"DB_PASSWORD"`
	DBSOURCE         string `mapstructure:"DB_SOURCE"`
	SecretKey        string `mapstructure:"SECRET_KEY"`
	SMTPUSERNAME     string `mapstructure:"SMTPUSERNAME"`
	SMTPHTTPPASSWORD string `mapstructure:"SMTPHTTPPASSWORD"`
	SMTPHOST         string `mapstructure:"SMTPHOST"`
	SMTPPORT         string `mapstructure:"SMTPPORT"`
}

var envs = []string{
	"DB_HOST", "DB_NAME", "DB_USER", "DB_PORT", "DB_PASSWORD", "DB_SOURCE", "SECRET_KEY",
	"SMTPUSERNAME","SMTPHTTPPASSWORD","SMTPHOST","SMTPPORT",
}

func LoadConfig() (Config, error) {
	var config Config

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return config, err
		}

	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err

	}

	if err := validator.New().Struct(&config); err != nil {
		return config, err
	}
	return config, nil
}
