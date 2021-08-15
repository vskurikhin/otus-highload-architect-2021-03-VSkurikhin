package config

import (
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

type (
	// Config обеспечивает конфигурацию системы.
	Config struct {
		Cache    Cache
		DataBase DataBase
		Server   Server
		Logging  Logging
	}
	// DataBase обеспечивает конфигурацию подключения к БД.
	Cache struct {
		Host     string `envconfig:"CACHE_HOST" default:"localhost"`
		Port     int    `envconfig:"CACHE_PORT" default:"3301"`
		Username string `envconfig:"CACHE_USERNAME" default:"guest"`
	}
	// DataBase обеспечивает конфигурацию подключения к БД.
	DataBase struct {
		HostRw   string `envconfig:"DB_HOST_RW" default:"localhost"`
		PortRw   int    `envconfig:"DB_PORT_RW" default:"6033"`
		HostRo   string `envconfig:"DB_HOST_RO" default:"localhost"`
		PortRo   int    `envconfig:"DB_PORT_RO" default:"6033"`
		DBName   string `envconfig:"DB_NAME" default:"hl"`
		Username string `envconfig:"DB_USERNAME" default:"hl"`
		Password string `envconfig:"DB_PASSWORD" default:"password"`
	}
	// Logging предоставляет конфигурацию журналирования.
	Logging struct {
		Debug  bool `envconfig:"LOGS_DEBUG"`
		Trace  bool `envconfig:"LOGS_TRACE"`
		Color  bool `envconfig:"LOGS_COLOR"`
		Pretty bool `envconfig:"LOGS_PRETTY"`
		Text   bool `envconfig:"LOGS_TEXT"`
	}
	// Server предоставляет конфигурацию сервера.
	Server struct {
		Host       string `envconfig:"SERVER_HOST" default:"0.0.0.0"`
		Port       string `envconfig:"PORT" default:"8080"`
		JWTSignKey string `envconfig:"JWT_SIGN_KEY" default:"TestForFastHTTPWithJWT"`
	}
)

// Environ возвращает настройки из окружения.
func Environ() (*Config, error) {
	cfg := Config{}
	err := envconfig.Process("", &cfg)
	return &cfg, err
}

// String возвращает конфигурацию в строковом формате.
func (c *Config) String() string {
	out, _ := yaml.Marshal(c)
	return string(out)
}
