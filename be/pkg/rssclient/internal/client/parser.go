package client

import (
	"context"
	"log/slog"
	"strings"
	"time"

	"github.com/ggrrrr/rss-viewer-task/be/pkg/rssclient"
)

func (r *rssRoot) parseRSS(ctx context.Context) []*rssclient.RssItem {
	// TODO add OTEL span
	rssItems := make([]*rssclient.RssItem, 0, len(r.ItemList))
	for i := range r.ItemList {
		item := r.createRssItem()
		err := r.ItemList[i].updateDetails(item)
		if err != nil {
			// TODO add span err and logging
			slog.ErrorContext(ctx, "rss parse", slog.Attr{Key: "error", Value: slog.AnyValue(err)})
		}
		rssItems = append(rssItems, item)
	}
	return rssItems
}

func (r *rssRoot) createRssItem() *rssclient.RssItem {
	rssItem := rssclient.RssItem{}
	rssItem.Source = r.ChannelTitle
	rssItem.SourceURL = r.ChannelLink
	return &rssItem
}

func (fromItem *rssItem) updateDetails(toItem *rssclient.RssItem) error {
	toItem.Title = fromItem.Title
	toItem.Description = fromItem.Description
	if fromItem.Link != "" {
		toItem.Link = fromItem.Link
	}
	pubDate, err := parseTime(fromItem.PubDate)
	if err != nil {
		return err
	}
	toItem.PublishDate = pubDate
	return nil
}

var pubDateFormats []string = []string{
	"Mon, 2 January 2006 15:04 MST",
	"Mon, 2 January 2006, 15:04:05 MST",
	"Mon, 2 January 2006 15:04:05 MST",
	"Mon, 2 Jan 2006 15:04:05 MST",
	"Mon,02 Jan 2006 15:04:05 -0700",
	"Mon, 02 Jan 2006 15:04:05 -0700",
	"Mon, 02 Jan 2006 15:04:05 --0700",
}

func parseTime(rssTime string) (time.Time, error) {
	cleanTime := strings.TrimSpace(rssTime)
	if cleanTime == "" {
		return time.Time{}, ErrTimeEmpy
	}

	for _, tFormat := range pubDateFormats {
		if t1, err := time.Parse(tFormat, cleanTime); err == nil {
			return t1.UTC(), nil
		}
	}
	return time.Time{}, ErrUnsupportedTime
}
