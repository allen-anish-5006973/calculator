.PHONY: default

SHELL :=/bin/bash
GOPATH=$(go env gopath)

default: setup test

setup:
	go get -d github.com/axw/gocov/gocov
	go get -d github.com/t-yuki/gocover-cobertura
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH)/bin

lint:
	golangci-lint run

test:
	mkdir -p coverage/
	GO111MODULE=on go clean -testcache ./... && go test -race ./... -covermode=atomic -coverprofile=coverage/coverage.out

test.cover: test
	GO111MODULE=on gocov convert coverage/coverage.out | gocov report

test.report: test.cover
	GO111MODULE=on go tool cover -html coverage/coverage.out -o coverage/coverage.html
	GO111MODULE=on gocover-cobertura < coverage/coverage.out > coverage/coverage.xml

test.quality: lint
