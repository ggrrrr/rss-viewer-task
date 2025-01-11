package client

import "fmt"

var (
	ErrTimeEmpy        = fmt.Errorf("rss client time is ampty")
	ErrUnsupportedTime = fmt.Errorf("rss client time unsupported format")

	ErrHttpGet = fmt.Errorf("rss client ")

	ErrHttpBadRequest = fmt.Errorf("rss client bad request")
	ErrHttpNotFound   = fmt.Errorf("rss client feed not found")

	ErrHttpUnauthorized = fmt.Errorf("rss client unauthorized")

	ErrHttpSystem = fmt.Errorf("rss client system error")
)
