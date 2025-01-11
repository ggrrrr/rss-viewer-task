package rssclient

import (
	"context"
	"encoding/xml"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/ggrrrr/rss-viewer-task/be/pkg/rssclient/internal/client"
)

func TestParseRSS(t *testing.T) {

	tests := []struct {
		name    string
		fromRss client.RSSRoot
		result  []RssItem
	}{
		{
			name: "ok",
			fromRss: client.RSSRoot{
				XMLName:            xml.Name{},
				Version:            "2.0",
				ChannelTitle:       "channel title",
				ChannelLink:        "channel link",
				ChannelDescription: "channel description",
				ItemList: []client.RSSItem{
					{
						Title:       "item 1 title",
						Link:        "item 1 link",
						PubDate:     "Tue, 23 May 2019 02:15:00 -0700",
						Description: "item 1 description",
					},
				},
			},
			result: []RssItem{
				{
					Title:       "item 1 title",
					Source:      "channel title",
					SourceURL:   "channel link",
					Link:        "item 1 link",
					PublishDate: time.Date(2019, 5, 23, 9, 15, 0, 0, time.UTC),
					Description: "item 1 description",
				},
			},
		},
		{
			name: "time error",
			fromRss: client.RSSRoot{
				XMLName:            xml.Name{},
				Version:            "2.0",
				ChannelTitle:       "channel title",
				ChannelLink:        "channel link",
				ChannelDescription: "channel description",
				ItemList: []client.RSSItem{
					{
						Title:       "item 1 title",
						Link:        "item 1 link",
						PubDate:     "",
						Description: "item 1 description",
					},
				},
			},
			result: []RssItem{
				{
					Title:     "item 1 title",
					Source:    "channel title",
					SourceURL: "channel link",
					Link:      "item 1 link",
					// PublishDate: time.Date(2019, 5, 23, 9, 15, 0, 0, time.UTC),
					Description: "item 1 description",
				},
			},
		},
	}

	ctx := context.TODO()

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := parseRSS(ctx, tc.fromRss)
			require.Equal(t, tc.result, result)
		})
	}

}

func TestTime(t *testing.T) {
	tests := []struct {
		name string
		from string
		to   time.Time
		err  error
	}{
		{
			name: "ok format",
			from: "Tue, 23 May 2019 02:15:00 -0700",
			to:   time.Date(2019, 5, 23, 9, 15, 0, 0, time.UTC),
			err:  nil,
		},
		{
			name: "ok format with spaces",
			from: "\t  Tue, 23 May 2019 02:15:00 -0700  ",
			to:   time.Date(2019, 5, 23, 9, 15, 0, 0, time.UTC),
			err:  nil,
		},
		{
			name: "err empty from",
			from: "",
			// to:   time.Date(2019, 5, 23, 9, 15, 0, 0, time.UTC),
			err: ErrTimeEmpy,
		},
		{
			name: "err format",
			from: "asd bad format",
			// to:   time.Date(2019, 5, 23, 9, 15, 0, 0, time.UTC),
			err: ErrUnsupportedTime,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := parseTime(tc.from)
			if tc.err == nil {
				require.Equal(t, tc.to, actual)
			} else {
				require.Error(t, err)
				require.ErrorIs(t, err, tc.err)
			}
		})
	}
}
