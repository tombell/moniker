VERSION?=dev
COMMIT=$(shell git rev-parse HEAD | cut -c -8)

LDFLAGS=-ldflags "-X main.Version=$(VERSION) -X main.Commit=$(COMMIT)"
MODFLAGS=-mod=vendor
TESTFLAGS=-cover

PLATFORMS:=darwin linux windows

all: dev

dev:
	@echo building dist/moniker
	@go build ${MODFLAGS} ${LDFLAGS} -o dist/moniker ./cmd/moniker

prod: $(PLATFORMS)

$(PLATFORMS):
	@echo building dist/moniker-$@-amd64
	@GOOS=$@ GOARCH=amd64 go build ${MODFLAGS} ${LDFLAGS} -o dist/moniker-$@-amd64 ./cmd/moniker

test:
	go test ${MODFLAGS} ./...

clean:
	rm -fr dist

.PHONY: all dev prod $(PLATFORMS) test clean
