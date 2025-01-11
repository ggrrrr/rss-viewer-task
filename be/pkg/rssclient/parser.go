package rssclient

import (
	"log/slog"
	"strings"
	"time"

	"github.com/ggrrrr/rss-viewer-task/be/pkg/rssclient/internal/client"
)

func parseRSS(r client.RSSRoot) []RssItem {
	rssItems := make([]RssItem, 0, len(r.ItemList))
	for i := range r.ItemList {
		item := createRssItem(r)
		err := updateDetails(&r.ItemList[i], &item)
		if err != nil {
			// TODO add span err and logging
			slog.Error("rss parse", slog.String("error", err.Error()))
		}
		rssItems = append(rssItems, item)
	}
	return rssItems
}

func createRssItem(r client.RSSRoot) RssItem {
	rssItem := RssItem{}
	rssItem.Source = r.ChannelTitle
	rssItem.SourceURL = r.ChannelLink
	return rssItem
}

func updateDetails(fromItem *client.RSSItem, toItem *RssItem) error {
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
