package config

import (
	"flag"
	"fmt"

	"github.com/caarlos0/env"
)

type (
	ServerConfig struct {
		// ListenAddress - адрес сервера сбора метрик
		ListenAddress string `env:"ADDRESS"`
	}
)

func NewServerConfig() ServerConfig {
	cfg := ServerConfig{
		ListenAddress: "localhost:8080",
	}

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Использование:\n")
		flag.PrintDefaults()
	}

	if flag.Lookup("a") == nil {
		flag.StringVar(&cfg.ListenAddress, "a", cfg.ListenAddress, "адрес и порт сервера сбора метрик")
	}

	flag.Parse()

	_ = env.Parse(&cfg)

	return cfg
}
