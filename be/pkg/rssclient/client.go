package rssclient

import (
	"context"
	"log/slog"
	"sync"

	"github.com/ggrrrr/rss-viewer-task/be/pkg/rssclient/internal/client"
)

func Parse(ctx context.Context, urls []string) []RssItem {
	// TODO add OTEL span

	var wg sync.WaitGroup
	var result = make([]RssItem, 0)
	var ch = make(chan []RssItem)

	wg.Add(len(urls))
	for _, url := range urls {
		go func() {
			defer wg.Done()
			result, err := client.Fetch(ctx, url)
			if err != nil {
				slog.ErrorContext(ctx, "client.Fetch", slog.String("url", url), slog.String("error", err.Error()))
				return
			}
			ch <- parseRSS(ctx, result)
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
		slog.InfoContext(ctx, "wg.Wait.")
	}()

	for items := range ch {
		result = append(result, items...)
	}

	return result
}
