package rest

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/ggrrrr/rss-viewer-task/be/pkg/common/auth"
	"github.com/ggrrrr/rss-viewer-task/be/svc/rssaggregator/intternal/app"
)

type MockApp struct {
	mock.Mock
}

var _ (application) = (*MockApp)(nil)

func (m *MockApp) Fetch(ctx context.Context, urls []string) []app.RSSItem {
	args := m.Called(urls[0])
	return args.Get(0).([]app.RSSItem)
}

func TestFetch(t *testing.T) {
	mockApp := new(MockApp)

	tests := []struct {
		name               string
		authInfo           auth.AuthInfo
		requestBody        string
		prepFunc           func(t *testing.T)
		responseStatusCode int
		responseBody       string
	}{
		{
			name:     "ok",
			authInfo: auth.AuthInfo{User: "admin"},
			prepFunc: func(t *testing.T) {
				mockApp.On("Fetch", "url1").Return([]app.RSSItem{
					{
						Title:       "title",
						PublishDate: time.Date(2019, 1, 23, 1, 15, 0, 0, time.UTC),
					},
				}, nil)
			},
			requestBody:        `{"urls":["url1"]}`,
			responseStatusCode: 200,
			responseBody:       `{"items":[{"title":"title","publish_date":"2019-01-23T01:15:00Z"}]}`,
		},
		{
			name: "err 401",
			prepFunc: func(t *testing.T) {
			},
			requestBody:        `{"urls":["url1"]}`,
			responseStatusCode: 401,
			responseBody:       `{"code":401,"error":"401 unauthorized"}`,
		},
		{
			name:     "err 400",
			authInfo: auth.AuthInfo{User: "admin"},
			prepFunc: func(t *testing.T) {
			},
			requestBody:        "",
			responseStatusCode: 400,
			responseBody:       `{"code":400,"error":"http body is empty"}`,
		},
		{
			name:     "err 400 bad request json",
			authInfo: auth.AuthInfo{User: "admin"},
			prepFunc: func(t *testing.T) {
			},
			requestBody:        "{asd}",
			responseStatusCode: 400,
			responseBody:       `{"code":400,"error":"invalid character 'a' looking for beginning of object key string"}`,
		},
		{
			name:     "err 400 empty url list",
			authInfo: auth.AuthInfo{User: "admin"},
			prepFunc: func(t *testing.T) {
			},
			requestBody:        "{}",
			responseStatusCode: 400,
			responseBody:       `{"code":400,"error":"empty list of url"}`,
		},
	}

	testServer := server{
		app: mockApp,
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			testRecorder := httptest.NewRecorder()
			ctx := context.Background()
			if tc.authInfo.User != "" {
				ctx = auth.Inject(ctx, tc.authInfo)
			}
			httpReq := httptest.NewRequest(http.MethodPost, "/someurl", strings.NewReader(tc.requestBody))
			httpReqWithAuth := httpReq.WithContext(ctx)

			tc.prepFunc(t)

			testServer.fetchRSS(testRecorder, httpReqWithAuth)

			require.Equal(t, tc.responseStatusCode, testRecorder.Result().StatusCode)
			responeBytes, err := io.ReadAll(testRecorder.Body)

			require.NoError(t, err)
			require.Equal(t, tc.responseBody, string(responeBytes))
		})
	}

}
