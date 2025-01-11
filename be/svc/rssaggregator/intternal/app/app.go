package app

import (
	"context"
	"log/slog"
	"time"

	"github.com/ggrrrr/rss-viewer-task/be/pkg/rssclient"
)

type (
	parser interface {
		Parse(ctx context.Context, urls []string) []rssclient.RssItem
	}

	App struct {
		parser parser
	}
)

func New(p parser) App {
	return App{
		parser: p,
	}
}

type RSSItem struct {
	Title       string    `json:"title,omitempty"`
	Source      string    `json:"source,omitempty"`
	SourceURL   string    `json:"source_url,omitempty"`
	Link        string    `json:"link,omitempty"`
	Descr       string    `json:"description,omitempty"`
	PublishDate time.Time `json:"publish_date,omitempty"`
}

func convertToResponse(from rssclient.RssItem) RSSItem {
	return RSSItem{
		Title:       from.Title,
		Source:      from.Source,
		SourceURL:   from.SourceURL,
		Link:        from.Link,
		Descr:       from.Description,
		PublishDate: from.PublishDate,
	}
}

func (a App) Fetch(ctx context.Context, urls []string) []RSSItem {
	slog.Info("app.Fetch", slog.Any("urls", urls))
	results := a.parser.Parse(ctx, urls)
	out := make([]RSSItem, len(results))
	for i := range results {
		out[i] = convertToResponse(results[i])
	}
	return out
}
