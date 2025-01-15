package system

import (
	"context"
	"log/slog"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"golang.org/x/sync/errgroup"

	"github.com/ggrrrr/rss-viewer-task/be/pkg/common/auth"
)

func (s *System) startWEB(ctx context.Context) {
	slog.Info("startWEB")

	ctx, cancel := context.WithCancel(ctx)
	webServer := &http.Server{
		Addr:    s.cfg.ListenAddr,
		Handler: s.mux,
	}

	group, gCtx := errgroup.WithContext(ctx)
	group.Go(func() error {
		slog.Info("rest starting...",
			slog.String("ListenAddr", s.cfg.ListenAddr),
		)

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

		return nil
	})

	s.shutdownFunc = append(s.shutdownFunc, func() error {
		slog.Info("web.shutdown")
		cancel()
		webServer.Close()
		return nil
	})

	s.startupFunc = append(s.startupFunc, func() error {
		return group.Wait()
	})
}

func initHTTPRouter(s *System) {
	s.mux = chi.NewRouter()
	s.mux.Use(s.httpHandlerCORS)
	s.mux.Use(middleware.Heartbeat("/liveness"))
	s.mux.Use(middleware.Logger)
	s.mux.Use(s.httpHandlerAuth)
	s.mux.Use(middleware.Recoverer)

	s.startWEB(context.Background())
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

// TODO unit testt
func (s *System) httpHandlerAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			next.ServeHTTP(w, r)
		}()

		authHeader := r.Header.Get("Authorization")
		if len(authHeader) == 0 {
			slog.ErrorContext(r.Context(), "httpHandlerAuth.NoHeader")
			return
		}

		splitToken := strings.Split(authHeader, " ")
		if len(splitToken) != 2 {
			slog.ErrorContext(r.Context(), "httpHandlerAuth.NoHeader.Bearer")
			return
		}
		slog.InfoContext(r.Context(), "httpHandlerAuth", slog.Any("Authorization", splitToken[1]))

		authInfo, err := s.verifier.Verify(splitToken[1])
		if err != nil {
			slog.ErrorContext(r.Context(), "httpHandlerAuth", slog.Any("error", err))
			return
		}

		slog.InfoContext(r.Context(), "httpHandlerAuth", slog.Any("authInfo", authInfo))
		ctx := auth.Inject(r.Context(), authInfo)
		r = r.WithContext(ctx)
	})
}
