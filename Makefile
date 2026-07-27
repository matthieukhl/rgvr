VERSION := $(shell git describe --tags --always)
GO_VERSION := $(shell go version | awk '{print $$3}')

LDFLAGS := -X github.com/matthieukhl/rgvr/cmd.appVersion=$(VERSION) -X github.com/matthieukhl/rgvr/cmd.goVersion=$(GO_VERSION)

build:
	go build -ldflags "$(LDFLAGS)" -o rgvr main.go