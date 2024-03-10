package app

import (
	"fmt"
	"net/http"

	"github.com/a-x-a/go-metric/internal/handler"
	"github.com/a-x-a/go-metric/internal/service/metricservice"
	"github.com/a-x-a/go-metric/internal/storage"
)

type (
	Server interface {
		Run() error
	}

	serverConfig struct {
		// ListenAddress - адрес сервера сбора метрик
		ListenAddress string
	}

	server struct {
		Config  serverConfig
		Storage storage.Storage
	}
)

func newServerConfig() serverConfig {
	return serverConfig{
		ListenAddress: "localhost:8080",
	}
}

func NewServer() *server {
	return &server{
		Config:  newServerConfig(),
		Storage: storage.NewMemStorage(),
	}
}

func (s *server) Run() error {
	service := metricservice.New(s.Storage)
	r := handler.Router(service)

	fmt.Println("listening on", s.Config.ListenAddress)

	err := http.ListenAndServe(s.Config.ListenAddress, r)

	if err != nil {
		return err
	}

	return nil
}
