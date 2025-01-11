FROM golang:1.23-alpine AS builder

WORKDIR /build

COPY ./go.work ./go.work.sum ./
COPY ./be ./be

RUN go build \
    -o app \  
    be/svc/rssaggregator/cmd/main.go

FROM alpine:latest AS runtime

WORKDIR /app

RUN addgroup -g 1000 -S app && \
    adduser -u 1000 -h /app -G app -S app

COPY --from=builder --chown=app:app /build/app /app/app

