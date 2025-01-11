package rssclient

import "fmt"

var (
	ErrTimeEmpy        = fmt.Errorf("rss client time is ampty")
	ErrUnsupportedTime = fmt.Errorf("rss client time unsupported format")
)
