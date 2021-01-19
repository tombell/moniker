VERSION?=dev
COMMIT=$(shell git rev-parse HEAD | cut -c -8)

LDFLAGS=-ldflags "-X main.version=${VERSION} -X main.commit=${COMMIT}"
MODFLAGS=-mod=vendor

PLATFORMS:=darwin linux windows

all: dev

clean:
	rm -fr dist/

dev:
	go build ${MODFLAGS} ${LDFLAGS} -o dist/moniker ./cmd/moniker

dist: $(PLATFORMS)

$(PLATFORMS):
	GOOS=$@ GOARCH=amd64 go build ${MODFLAGS} ${LDFLAGS} -o dist/moniker-$@-amd64 ./cmd/moniker

test:
	go test ${MODFLAGS} ./...

modules:
	@go get -u ./... && go mod download && go mod tidy && go mod vendor

.PHONY: all clean dev dist $(PLATFORMS) test modules
