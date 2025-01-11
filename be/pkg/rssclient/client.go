package rssclient

import (
	"context"
	"log/slog"
	"sync"

	"github.com/ggrrrr/rss-viewer-task/be/pkg/rssclient/internal/client"
)

type RSSParser struct {
}

func New() RSSParser {
	return RSSParser{}
}

// This is to help us with unit testing
func (RSSParser) Parse(ctx context.Context, urls []string) []RssItem {
	var wg sync.WaitGroup
	var result = make([]RssItem, 0)
	var ch = make(chan []RssItem)

	wg.Add(len(urls))
	for _, url := range urls {
		go func() {
			defer wg.Done()
			result, err := client.Fetch(ctx, url)
			if err != nil {
				slog.Error("rssclient.client.Fetch", slog.String("url", url), slog.String("error", err.Error()))
				return
			}
			ch <- parseRSS(result)
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
		slog.Debug("rssclient.Parse.wg.Wait.")
	}()

	for items := range ch {
		result = append(result, items...)
	}

	return result
}
