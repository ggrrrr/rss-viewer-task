package rest

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/ggrrrr/rss-viewer-task/be/svc/rssaggregator/intternal/app"
)

type (
	application interface {
		Fetch(ctx context.Context, urls []string) []app.RSSItem
	}

	server struct {
		app application
	}
)

func Router(a application) http.Handler {
	s := server{
		app: a,
	}
	router := chi.NewRouter()
	router.Post("/parse", s.fetchRSS)
	return router
}
