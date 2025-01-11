package client

import (
	"context"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/ggrrrr/rss-viewer-task/be/pkg/rssclient"
)

const supportedRSSVersion string = "2.0"

func Fetch(ctx context.Context, url string) ([]*rssclient.RssItem, error) {
	// TODO add OTEL Span
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("cant call url %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode == 400 {
		return nil, ErrHttpBadRequest
	}
	if res.StatusCode == 404 {
		return nil, ErrHttpNotFound
	}

	if res.StatusCode == 401 ||
		res.StatusCode == 403 {
		return nil, ErrHttpUnauthorized
	}

	if res.StatusCode >= 500 {
		return nil, ErrHttpSystem
	}

	decoder := xml.NewDecoder(res.Body)
	root := rssRoot{}

	err = decoder.Decode(&root)
	if err != nil {
		return nil, fmt.Errorf("cant parse data %w", err)
	}

	if root.Version != supportedRSSVersion {
		return nil, fmt.Errorf("unsupported version %s", root.Version)
	}

	// Here we dont really care if we have item parsing errors
	// We will only log these
	result := root.parseRSS(ctx)

	fmt.Printf("\n%+v", result)
	return result, nil
}
