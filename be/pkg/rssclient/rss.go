package rssclient

import "time"

type RssItem struct {
	// Title: from parsed RSS Item -> rss.channel.item[].Title
	Title string `json:"title,omitempty"`
	// Description: from parsed RSS channel -> rss.channel.title
	Source string `json:"source,omitempty"`
	// SourceURL: parsed from RSS channel -> rss.channel.link
	SourceURL string `json:"source_url,omitempty"`
	// Link ( URL ) which is part of the -> rss.channel.item[].link
	Link string `json:"link,omitempty"`
	// Published date of the RSS Item -> rss.channel.item[].pubDate converted to UTC
	PublishDate time.Time `json:"publish_date,omitempty"`
	// Description of the RSS Item -> rss.channel.item[].Description
	Description string `json:"description,omitempty"`
}
