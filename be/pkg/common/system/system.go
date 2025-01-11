package system

import (
	"net/http"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/go-chi/chi"
)

type (
	config struct {
		Timeout    time.Duration `env:"HTTP_TTL" envDefault:"10s"`
		ListenAddr string        `env:"LISTEN_ADDR" envDefault:":8080"`
		CORSHosts  string        `env:"CORS_HOSTS"`
	}

	System struct {
		cfg config
		mux *chi.Mux
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
		cfg: cfg,
	}
	newHTTPRouter(system)

	return system, nil
}

func (s *System) MountAPI(pattern string, handler http.Handler) {
	s.mux.Mount(pattern, handler)
}
