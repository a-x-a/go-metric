package app

import (
	"context"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/a-x-a/go-metric/internal/config"
	"github.com/a-x-a/go-metric/internal/storage"
)

func TestNewServer(t *testing.T) {
	t.Run("create new server", func(t *testing.T) {
		got := NewServer()
		require.NotNil(t, got)
	})
}

func Test_serverRun(t *testing.T) {
	stor := storage.NewMemStorage()
	cfg := config.NewServerConfig()
	srv := server{
		Config:  cfg,
		Storage: stor,
		srv:     &http.Server{Addr: cfg.ListenAddress},
	}
	// ctx := context.Background()
	ctx, cancel := context.WithCancel(context.Background())
	time.AfterFunc(time.Second*10, cancel)
	defer cancel()
	// wg := sync.WaitGroup{}
	// wg.Add(1)
	go func() {
		// defer wg.Done()
		srv.Run(ctx)
	}()

	conn, err := net.Dial("tcp", srv.srv.Addr)
	require.NoError(t, err)
	defer conn.Close()
	require.NotNil(t, conn)

	_ = srv.srv.Shutdown(ctx)
	// if err := srv.srv.Shutdown(ctx); err != nil {
	// 	// panic(err) // failure/timeout shutting down the server gracefully
	// }

	// wg.Wait()
}
