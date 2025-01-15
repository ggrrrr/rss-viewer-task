package system

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/go-chi/chi"
	"golang.org/x/sync/errgroup"

	"github.com/ggrrrr/rss-viewer-task/be/pkg/common/auth"
)

type (
	config struct {
		CrtKeyFile string        `env:"CRT_KEY_FILE"`
		Timeout    time.Duration `env:"HTTP_TTL" envDefault:"10s"`
		ListenAddr string        `env:"LISTEN_ADDR" envDefault:":8080"`
		CORSHosts  string        `env:"CORS_HOSTS"`
	}

	System struct {
		cfg config
		mux *chi.Mux

		verifier auth.Verifier

		shutdownFunc []func() error
		startupFunc  []func() error
	}
)

func New() (*System, error) {
	var err error
	cfg := config{}

	err = env.Parse(&cfg)
	if err != nil {
		return nil, err
	}

	if cfg.ListenAddr == "" {
		panic("empty LISTEN_ADDR")
	}

	system := &System{
		cfg:          cfg,
		shutdownFunc: make([]func() error, 0, 1),
		startupFunc:  make([]func() error, 0, 1),
	}

	err = initJWT(system)
	if err != nil {
		return nil, err
	}

	initHTTPRouter(system)

	return system, nil
}

func (s *System) MountAPI(pattern string, handler http.Handler) {
	s.mux.Mount(pattern, handler)
}

func (s *System) Start(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)

	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	// }

	g.Go(func() error {
		<-ctx.Done()
		cancel()
		slog.Info("git kill signal")

		for _, shutdownFunc := range s.shutdownFunc {
			err := shutdownFunc()
			if err != nil {
				return err
			}
		}

		return nil
	})

	g.Go(func() error {
		for _, startFunc := range s.startupFunc {
			err := startFunc()
			if err != nil {
				return err
			}
		}
		return nil
	})

	return g.Wait()
}
