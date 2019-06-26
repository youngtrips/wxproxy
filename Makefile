## env
CGO_ENABLED	:= 0
GOARCH		:= amd64
GOOS		:= $(shell uname -s | tr 'A-Z' 'a-z')
GO			:= go

## version
VERSION 	:= 1.0.0

## targets
APPS=$(shell ls cmd)

all: build

build:
	@for APP in $(APPS) ; do \
		echo building $$APP ; \
		CGO_ENABLED=$(CGO_ENABLED) $(GO) build -ldflags "-X main.APP_VERSION=$(VERSION)" -o bin/$$APP ./cmd/$$APP; \
	done

clean:
	rm -rf bin
