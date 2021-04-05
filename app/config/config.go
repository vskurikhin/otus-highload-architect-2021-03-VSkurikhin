package config

import (
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

type (
	// Config обеспечивает конфигурацию системы.
	Config struct {
		Server  Server
		Logging Logging
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
		Host string `envconfig:"SERVER_HOST" default:"localhost:6060"`
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
