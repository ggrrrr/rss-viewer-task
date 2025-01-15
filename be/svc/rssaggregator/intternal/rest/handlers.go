package rest

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/ggrrrr/rss-viewer-task/be/pkg/common/auth"
	"github.com/ggrrrr/rss-viewer-task/be/pkg/common/web"
	"github.com/ggrrrr/rss-viewer-task/be/svc/rssaggregator/intternal/app"
)

type (
	fetchRSSResponse struct {
		Items []app.RSSItem `json:"items"`
	}

	fetchRSSRequest struct {
		URL []string `json:"urls"`
	}
)

func (s server) fetchRSS(w http.ResponseWriter, r *http.Request) {
	user := auth.Extract(r.Context())
	if !auth.HasAccess(user) {
		slog.Warn("fetch.Unauthorized", slog.Any("user", user))
		web.SendUnauthorized(w)
		return
	}

	if r.ContentLength == 0 {
		web.SendBadRequest(w, fmt.Errorf("http body is empty"))
		return
	}

	if r.Body == nil {
		web.SendBadRequest(w, fmt.Errorf("http body is nil"))
		return
	}
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	fetchRequest := fetchRSSRequest{}
	err := decoder.Decode(&fetchRequest)
	if err != nil {
		slog.Error("parse error", slog.Any("error", err))
		web.SendBadRequest(w, err)
		return
	}

	if len(fetchRequest.URL) == 0 {
		slog.Error("no result")
		web.SendBadRequest(w, fmt.Errorf("empty list of url"))
		return
	}

	result := s.app.Fetch(r.Context(), fetchRequest.URL)
	web.SendPayload(w, fetchRSSResponse{
		Items: result,
	})
}
