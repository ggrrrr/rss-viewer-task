package app

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/ggrrrr/rss-viewer-task/be/pkg/rssclient"
)

type MockParser struct {
	mock.Mock
}

func (m *MockParser) Parse(ctx context.Context, urls []string) []rssclient.RssItem {
	args := m.Called(urls[0])
	return args.Get(0).([]rssclient.RssItem)
}

var _ (parser) = (*MockParser)(nil)

func TestFetch(t *testing.T) {

	testParser := new(MockParser)

	tests := []struct {
		name     string
		urls     []string
		prepFunc func(t *testing.T)
		result   []RSSItem
	}{
		{
			name: "ok one url",
			urls: []string{"url1"},
			prepFunc: func(t *testing.T) {
				testParser.On("Parse", "url1").
					Return(
						[]rssclient.RssItem{
							{
								Title: "title",
							},
						},
					)
			},
			result: []RSSItem{
				{
					Title: "title",
				},
			},
		},
		{
			name: "ok no url",
			urls: []string{"url2"},
			prepFunc: func(t *testing.T) {
				testParser.On("Parse", "url2").
					Return(
						[]rssclient.RssItem{},
					)
			},
			result: []RSSItem{},
		},
	}

	testApp := App{
		parser: testParser,
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.prepFunc(t)
			result := testApp.Fetch(context.TODO(), tc.urls)
			require.Equal(t, tc.result, result)
		})
	}
}

func TestConvert(t *testing.T) {
	tests := []struct {
		name string
		from rssclient.RssItem
		to   RSSItem
	}{
		{
			name: "ok",
			from: rssclient.RssItem{
				Title:       "tittle",
				Source:      "source",
				SourceURL:   "source url",
				Link:        "link",
				PublishDate: time.Date(2019, 2, 1, 1, 1, 0, 0, time.UTC),
			},
			to: RSSItem{
				Title:       "tittle",
				Source:      "source",
				SourceURL:   "source url",
				Link:        "link",
				PublishDate: time.Date(2019, 2, 1, 1, 1, 0, 0, time.UTC),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := convertToResponse(tc.from)
			require.Equal(t, tc.to, actual)
		})
	}
}
