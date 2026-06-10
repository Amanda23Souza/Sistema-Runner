.PHONY: build test clean

VERSION ?= $(shell git describe --tags --always 2>/dev/null || echo "v0.0.0-dev")
COMMIT ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
BUILDTIME ?= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ" 2>/dev/null || echo "unknown")

LDFLAGS = -s -w \
	-X 'github.com/Amanda23Souza/Sistema-Runner/cli-assinatura/internal/version.Version=$(VERSION)' \
	-X 'github.com/Amanda23Souza/Sistema-Runner/cli-assinatura/internal/version.Commit=$(COMMIT)' \
	-X 'github.com/Amanda23Souza/Sistema-Runner/cli-assinatura/internal/version.BuildTime=$(BUILDTIME)'

build:
	cd cli-assinatura && go build -ldflags="$(LDFLAGS)" -o assinatura.exe ./cmd/assinatura

test:
	cd cli-assinatura && go test -v -race ./...

clean:
	rm -f cli-assinatura/assinatura.exe
