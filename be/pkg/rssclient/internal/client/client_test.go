package client

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/ggrrrr/rss-viewer-task/be/pkg/rssclient"
	"github.com/ggrrrr/rss-viewer-task/be/pkg/rssclient/testdata"
	// "github.com/ggrrrr/rss-viewer-task/be/testdata"
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
			w.Write(content)
		}
	}))

}

func TestFetch(t *testing.T) {

	tests := []struct {
		name   string
		server *httptest.Server
		result []*rssclient.RssItem
		err    error
	}{
		{
			name:   "ok",
			server: createTestResponse(t, 200, "small_feed.xml"),
			result: []*rssclient.RssItem{
				{
					Title:       "item 1 title",
					Source:      "Channel Title",
					SourceURL:   "https://ress.serveri/feed/",
					Link:        "http://itemlink/",
					PublishDate: time.Date(2019, 1, 8, 1, 15, 0, 0, time.UTC),
					Description: "item 1 description",
				},
			},
			err: nil,
		},
		{
			name:   "error 400",
			server: createTestResponse(t, 400, ""),
			result: []*rssclient.RssItem{},
			err:    ErrHttpBadRequest,
		},
		{
			name:   "error 401",
			server: createTestResponse(t, 401, ""),
			result: []*rssclient.RssItem{},
			err:    ErrHttpUnauthorized,
		},
		{
			name:   "error 404",
			server: createTestResponse(t, 404, ""),
			result: []*rssclient.RssItem{},
			err:    ErrHttpNotFound,
		},
		{
			name:   "error 500",
			server: createTestResponse(t, 500, ""),
			result: []*rssclient.RssItem{},
			err:    ErrHttpSystem,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			defer tc.server.Close()

			result, err := Fetch(context.TODO(), tc.server.URL)
			if tc.err == nil {
				require.NoError(t, err)
				require.Equal(t, tc.result, result)
			} else {
				require.Error(t, err)
				require.ErrorIs(t, err, tc.err)
			}
		})
	}
}
