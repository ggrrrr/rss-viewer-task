package client

import (
	"encoding/xml"
)

type rssRoot struct {
	XMLName            xml.Name  `xml:"rss"`
	Version            string    `xml:"version,attr"`
	ChannelTitle       string    `xml:"channel>title"`
	ChannelLink        string    `xml:"channel>link"`
	ChannelDescription string    `xml:"channel>description"`
	ChannelPubDate     string    `xml:"channel>pubDate"`
	ItemList           []rssItem `xml:"channel>item"`
}

type rssItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	PubDate     string `xml:"pubDate"`
	Description string `xml:"description"`
}
