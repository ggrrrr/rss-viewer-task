package rssclient

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/ggrrrr/rss-viewer-task/be/pkg/rssclient/testdata"
)

var pwd = testdata.RepoDir()

func createTestResponse(t *testing.T, status int, file string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		if file != "" {
			content, err := os.ReadFile(fmt.Sprintf("%s/rssclient/testdata/rssfeed/%s", pwd, file))
			if err != nil {
				t.Fatal(t, err, "read test file error")
			}
			_, err = w.Write(content)
			if err != nil {
				t.Fatal(t, err, "write http body error")
			}
		}
	}))

}

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		servers []*httptest.Server
		result  []RssItem
	}{
		{
			name: "ok",
			servers: []*httptest.Server{
				createTestResponse(t, 200, "small_feed.xml"),
				// createTestResponse(t, 200, "small_feed_2.xml"),
				// createTestResponse(t, 400, ""),
			},
			result: []RssItem{
				{
					Title:       "item 1 title",
					Source:      "Channel Title",
					SourceURL:   "https://ress.serveri/feed/",
					Link:        "http://itemlink/",
					PublishDate: time.Date(2019, 1, 8, 1, 15, 0, 0, time.UTC),
					Description: "item 1 description",
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			urls := make([]string, len(tc.servers))
			for i := range tc.servers {
				urls[i] = tc.servers[i].URL
			}
			result := Parse(context.TODO(), urls)
			require.Equal(t, tc.result, result)

		})
	}

}

func TestMultiURL(t *testing.T) {
	tests := []struct {
		name    string
		servers []*httptest.Server
	}{
		{
			name: "ok",
			servers: []*httptest.Server{
				createTestResponse(t, 200, "small_feed.xml"),
				createTestResponse(t, 200, "small_feed_2.xml"),
				createTestResponse(t, 200, "small_feed_2.xml"),
				createTestResponse(t, 400, ""),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			urls := make([]string, len(tc.servers))
			for i := range tc.servers {
				urls[i] = tc.servers[i].URL
			}
			result := Parse(context.TODO(), urls)
			require.Equal(t, len(result), 5)

		})
	}

}
