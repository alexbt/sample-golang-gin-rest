#!/bin/sh

go generate ./...
go fmt ./...
#env GOOS=linux GOARCH=amd64 go build ./cmd/... ./pkg/...
docker build -t mytest .
