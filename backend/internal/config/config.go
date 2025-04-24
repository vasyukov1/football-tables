package config

import "github.com/spf13/viper"

type Config struct {
	DB struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
	}
	HTTP struct {
		Port string
	}
	JWT struct {
		Secret string
	}
	CORS struct {
		AllowedOrigins []string
		Debug          bool
	}
}

func Load() *Config {
	v := viper.New()
	v.AutomaticEnv()

	v.SetDefault("HTTP_PORT", "8080")
	v.SetDefault("DB_PORT", "5432")

	return &Config{
		DB: struct {
			Host     string
			Port     string
			User     string
			Password string
			Name     string
		}{
			Host:     v.GetString("DB_HOST"),
			Port:     v.GetString("DB_PORT"),
			User:     v.GetString("DB_USER"),
			Password: v.GetString("DB_PASSWORD"),
			Name:     v.GetString("DB_NAME"),
		},
		HTTP: struct {
			Port string
		}{
			Port: v.GetString("HTTP_PORT"),
		},
		JWT: struct {
			Secret string
		}{
			Secret: v.GetString("JWT_SECRET"),
		},
		CORS: struct {
			AllowedOrigins []string
			Debug          bool
		}{
			AllowedOrigins: v.GetStringSlice("CORS_ALLOWED_ORIGINS"),
			Debug:          v.GetBool("CORS_DEBUG"),
		},
	}
}
