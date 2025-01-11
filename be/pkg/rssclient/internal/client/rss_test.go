package client

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRSSDecode(t *testing.T) {

	tests := []struct {
		name    string
		fromXml string
		result  rssRoot
		expErr  error
	}{
		{
			name: "ok v2",
			fromXml: `<?xml version='1.0' encoding='UTF-8'?>
<rss xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#" xmlns:taxo="http://purl.org/rss/1.0/modules/taxonomy/" xmlns:sy="http://purl.org/rss/1.0/modules/syndication/" xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:itunes="http://www.itunes.com/dtds/podcast-1.0.dtd" xmlns:slash="http://purl.org/rss/1.0/modules/slash/" version="2.0">
<channel>
<title>channel title</title>
<link>channel link</link>
<description>channel description</description>
<pubDate>Thu, 26 Feb 2015 14:15:27 GMT</pubDate>
<item>
<title>item 1 title</title>
<description>item 1 description</description>
<pubDate>Thu, 26 Feb 2015 14:14:59 GMT</pubDate>
</item>
</channel>
</rss>
  `,
			result: rssRoot{
				XMLName:            xml.Name{Space: "", Local: "rss"},
				Version:            "2.0",
				ChannelTitle:       "channel title",
				ChannelLink:        "channel link",
				ChannelDescription: "channel description",
				ChannelPubDate:     "Thu, 26 Feb 2015 14:15:27 GMT",
				ItemList: []rssItem{
					{
						Title:       "item 1 title",
						PubDate:     "Thu, 26 Feb 2015 14:14:59 GMT",
						Description: "item 1 description",
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			root := rssRoot{}
			err := xml.Unmarshal([]byte(tc.fromXml), &root)
			if tc.expErr == nil {
				require.Equal(t, tc.result, root)
			} else {
				// TODO Use ErrorAs or ErrorIs for more detailed testing
				require.Error(t, err)
			}
		})
	}

}
