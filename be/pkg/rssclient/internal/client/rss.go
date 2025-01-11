package client

import (
	"encoding/xml"
)

type RSSRoot struct {
	XMLName            xml.Name  `xml:"rss"`
	Version            string    `xml:"version,attr"`
	ChannelTitle       string    `xml:"channel>title"`
	ChannelLink        string    `xml:"channel>link"`
	ChannelDescription string    `xml:"channel>description"`
	ItemList           []RSSItem `xml:"channel>item"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	PubDate     string `xml:"pubDate"`
	Description string `xml:"description"`
}
