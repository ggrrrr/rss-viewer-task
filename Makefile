DOCKER_REPO ?= "local"
GIT_HASH ?= $(shell git log --format="%h" -n 1)

go_sync:
	go work sync

go_test:
	go test -cover  ./be/pkg/rssclient/...
	go test -cover  ./be/svc/rssaggregator/...

go_lint:
	golangci-lint run -v ./be/pkg/rssclient/...
	golangci-lint run -v ./be/pkg/common/...
	golangci-lint run -v ./be/svc/rssaggregator/...

build_svc:
	docker build \
		-f ./.docker/be.Dockerfile \
		--tag "${DOCKER_REPO}/svc/rssaggregator:${GIT_HASH}" \
		./
	
	docker tag ${DOCKER_REPO}/svc/rssaggregator:${GIT_HASH} ${DOCKER_REPO}/svc/rssaggregator:latest

