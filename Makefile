DOCKER_REPO ?= "local"
GIT_HASH ?= $(shell git log --format="%h" -n 1)

go_clean:
	cd be/pkg/common && go mod tidy
	cd be/svc/rssaggregator && go mod tidy
	# go clean -cache
	# go clean -testcache
	# go clean -fuzzcache
	# go clean -modcache

go_test:
	go test -cover  ./be/pkg/rssclient/...

go_lint:
	golangci-lint run -v ./be/pkg/rssclient/...

