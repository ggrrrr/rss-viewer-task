package client

import (
	"context"
	"encoding/xml"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

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
			_, err = w.Write(content)
			if err != nil {
				t.Fatal(t, err, "write http body error")
			}
		}
	}))

}

func TestFetch(t *testing.T) {

	tests := []struct {
		name   string
		server *httptest.Server
		result RSSRoot
		err    error
	}{
		{
			name:   "ok",
			server: createTestResponse(t, 200, "small_feed.xml"),
			result: RSSRoot{
				XMLName:            xml.Name{Space: "", Local: "rss"},
				Version:            "2.0",
				ChannelTitle:       "Channel Title",
				ChannelLink:        "https://ress.serveri/feed/",
				ChannelDescription: "description",
				ItemList: []RSSItem{
					{
						Title:       "item 1 title",
						Link:        "http://itemlink/",
						PubDate:     "Tue, 8 Jan 2019 01:15:00 GMT",
						Description: "item 1 description",
					},
				},
			},
			err: nil,
		},
		{
			name:   "error 400",
			server: createTestResponse(t, 400, ""),
			// result: []*rssclient.RssItem{},
			err: ErrHttpBadRequest,
		},
		{
			name:   "error 401",
			server: createTestResponse(t, 401, ""),
			// result: []*rssclient.RssItem{},
			err: ErrHttpUnauthorized,
		},
		{
			name:   "error 404",
			server: createTestResponse(t, 404, ""),
			// result: []*rssclient.RssItem{},
			err: ErrHttpNotFound,
		},
		{
			name:   "error 500",
			server: createTestResponse(t, 500, ""),
			// result: []*rssclient.RssItem{},
			err: ErrHttpSystem,
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
