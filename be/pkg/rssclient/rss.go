package rssclient

import "time"

type RssItem struct {
	// Title: from parsed RSS Item -> rss.channel.item[].Title
	Title string
	// Description: from parsed RSS channel -> rss.channel.title
	Source string
	// SourceURL: parsed from RSS channel -> rss.channel.link
	SourceURL string
	// Link ( URL ) which is part of the -> rss.channel.item[].link
	Link string
	// Published date of the RSS Item -> rss.channel.item[].pubDate converted to UTC
	PublishDate time.Time
	// Description of the RSS Item -> rss.channel.item[].Description
	Description string
}
