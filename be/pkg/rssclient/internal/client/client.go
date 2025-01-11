package client

import (
	"context"
	"encoding/xml"
	"fmt"
	"log/slog"
	"net/http"
)

const supportedRSSVersion string = "2.0"

func Fetch(ctx context.Context, url string) (root RSSRoot, err error) {
	// TODO add OTEL Span
	slog.DebugContext(ctx, "client.Parse", slog.String("url", url))
	res, err := http.Get(url)
	if err != nil {
		return root, fmt.Errorf("cant call url %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode == 400 {
		return root, ErrHttpBadRequest
	}
	if res.StatusCode == 404 {
		return root, ErrHttpNotFound
	}

	if res.StatusCode == 401 ||
		res.StatusCode == 403 {
		return root, ErrHttpUnauthorized
	}

	if res.StatusCode >= 500 {
		return root, ErrHttpSystem
	}

	decoder := xml.NewDecoder(res.Body)

	err = decoder.Decode(&root)
	if err != nil {
		return root, fmt.Errorf("cant parse data %w", err)
	}

	if root.Version != supportedRSSVersion {
		return root, fmt.Errorf("unsupported version %s", root.Version)
	}

	return root, nil
}
