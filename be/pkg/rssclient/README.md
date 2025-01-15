# Rss client

This is simple RSS Client.

* uses SLOG for logging
* OTEL for tracing //TODO

Entry Method: `Parse(ctx context.Context, urls []string)`

* `ctx` context which will be used to add any traceing data (OTEL)
* `urls` a array of string with URLs from which to fetch RSS data items
  * Any error during parsing  will be ignored, and only log will be produced

## Example use

### run the following

```sh
# get the package
go get github.com/ggrrrr/rss-viewer-task/be/pkg/rssclient

```

### Create simple main.go file

```golang
package main

import (
    "context"
    "fmt"

    "github.com/ggrrrr/rss-viewer-task/be/pkg/rssclient"
)

func main() {
    parser := rssclient.New() 
    items := parser.Parse(
        context.Background(),
        []string{"https://news.google.com/rss/search?hl=en-US&gl=US&q=samsung&um=1&ie=UTF-8&ceid=US:en"},
    )
    for _, item := range items {
        fmt.Printf(":%+v\n ", item)
    }
}
```

Aggregator for [RSS specification](https://www.rssboard.org/rss-specificationhttps://www.rssboard.org/rss-specification)

Example RSS providers: [google news](https://news.google.com/rss/search?hl=en-US&gl=US&q=samsung&um=1&ie=UTF-8&ceid=US:en)
