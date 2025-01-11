package system

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"golang.org/x/sync/errgroup"

	"github.com/ggrrrr/rss-viewer-task/be/pkg/common/auth"
)

func (s *System) StartWeb(ctx context.Context) error {
	slog.Info("StartWeb")
	// addr := s.cfg.Rest.Address
	webServer := &http.Server{
		Addr:    s.cfg.ListenAddr,
		Handler: s.mux,
	}

	group, gCtx := errgroup.WithContext(ctx)
	group.Go(func() error {
		slog.Info("rest starting...",
			slog.String("ListenAddr", s.cfg.ListenAddr),
		)
		defer slog.Info("web server shutdown")
		if err := webServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	})

	group.Go(func() error {
		<-gCtx.Done()
		slog.Info("rest shutdown...",
			slog.String("ListenAddr", s.cfg.ListenAddr),
		)
		ctx, cancel := context.WithTimeout(context.Background(), s.cfg.Timeout)
		defer cancel()
		if err := webServer.Shutdown(ctx); err != nil {
			return err
		}
		return group.Wait()
	})

	return group.Wait()
}

func newHTTPRouter(s *System) {
	s.mux = chi.NewRouter()
	// s.mux.NotFound(web.MethodNotFoundHandler)
	// s.mux.MethodNotAllowed(web.MethodNotAllowedHandler)
	s.mux.Use(s.httpHandlerCORS)
	s.mux.Use(middleware.Heartbeat("/liveness"))
	s.mux.Use(middleware.Logger)
	s.mux.Use(s.httpHandlerAuth)
	// s.mux.Use(middleware.Recoverer)

}

func (s *System) httpHandlerCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// READ FROM CFG
		defer next.ServeHTTP(w, r)
		if s.cfg.CORSHosts == "" {
			return
		}
		w.Header().Set("Access-Control-Allow-Origin", s.cfg.CORSHosts)
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Authorization")

		if r.Method == http.MethodOptions {
			out := "."
			w.WriteHeader(200)
			_, err := w.Write([]byte(out))
			if err != nil {
				slog.ErrorContext(r.Context(), "cant write body")
			}
			return
		}
	})
}

func (s *System) httpHandlerAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO READ  HEADER and parse JWT
		slog.InfoContext(r.Context(), "httpHandlerAuth")
		authInfo := auth.AuthInfo{
			User: "admin",
		}
		ctx := auth.Inject(r.Context(), authInfo)
		newReq := r.WithContext(ctx)
		next.ServeHTTP(w, newReq)
	})
}
