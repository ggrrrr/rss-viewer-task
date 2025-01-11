package app

import (
	"context"

	"github.com/ggrrrr/rss-viewer-task/be/pkg/rssclient"
)

type App struct {
}

type RSSItem struct {
	Title     string `json:"title"`
	Source    string `json:"source"`
	SourceURL string `json:"source_url"`
	Link      string `json:"link"`
	Descr     string `json:"description"`
}

func convertToResponse(from rssclient.RssItem) RSSItem {
	return RSSItem{
		Title:     from.Title,
		Source:    from.Source,
		SourceURL: from.SourceURL,
		Link:      from.Link,
		Descr:     from.Description,
	}
}

func (a App) Fetch(ctx context.Context, urls []string) []RSSItem {
	results := rssclient.Parse(ctx, urls)
	out := make([]RSSItem, len(results))
	for i := range results {
		out[i] = convertToResponse(results[i])
	}
	return out
}
