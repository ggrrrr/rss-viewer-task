package main

import (
	"context"

	"github.com/ggrrrr/rss-viewer-task/be/pkg/common/logger"
	"github.com/ggrrrr/rss-viewer-task/be/pkg/common/system"
	"github.com/ggrrrr/rss-viewer-task/be/pkg/rssclient"
	"github.com/ggrrrr/rss-viewer-task/be/svc/rssaggregator/intternal/app"
	"github.com/ggrrrr/rss-viewer-task/be/svc/rssaggregator/intternal/rest"
)

func main() {

	logger.Configure()

	s, err := system.New()
	if err != nil {
		panic(err)
	}

	a := app.New(rssclient.New())

	router := rest.Router(a)
	s.MountAPI("/v1", router)
	err = s.Start(context.Background())
	if err != nil {
		panic(err)
	}
}
