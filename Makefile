BINARY=moniker
GOARCH=amd64

VERSION?=dev
COMMIT=$(shell git rev-parse HEAD | cut -c -8)

LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Commit=${COMMIT}"

PACKAGE=./cmd/moniker

all: dev

clean:
	rm -fr dist/

dev:
	go build ${LDFLAGS} -o dist/${BINARY} ${PACKAGE}

cibuild:
	go build -mod=vendor ${LDFLAGS} -o dist/${BINARY} ${PACKAGE}

dist: linux darwin windows

linux:
	GOOS=linux GOARCH=${GOARCH} go build ${LDFLAGS} -o dist/${BINARY}-linux-${GOARCH} ${PACKAGE}

darwin:
	GOOS=darwin GOARCH=${GOARCH} go build ${LDFLAGS} -o dist/${BINARY}-darwin-${GOARCH} ${PACKAGE}

windows:
	GOOS=windows GOARCH=${GOARCH} go build ${LDFLAGS} -o dist/${BINARY}-windows-${GOARCH} ${PACKAGE}

test:
	@go test github.com/tombell/moniker

.PHONY: all clean dev cibuild dist linux darwin windows test
