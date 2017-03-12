.PHONY: clean version version-short fmt vet test

VERSION := 0.0.1
NAME := dialect
ORG := pinpt
PKG := $(ORG)/$(NAME)

SHELL := /bin/bash
BASEDIR := $(shell echo $${PWD})
BUILD := $(shell git rev-parse HEAD | cut -c1-8)
SRC := $(shell find . -type f -name '*.go' -not -path './vendor/*' -not -path './.git/*')
PKGMAIN := cmd/$(NAME)/main.go

L="-X=github.com/$(PKG)/cmd/main.Build=$(BUILD) -X=github.com/$(PKG)/cmd/main.Version=$(VERSION)"

all: version fmt vet test

version:
	@echo "version: $(VERSION) (build: $(BUILD))"

version-short:
	@echo $(VERSION)

clean:
	@rm -rf build recordings

fmt:
	@gofmt -s -l -w $(SRC)

vet:
	@for i in `find . -type d -not -path './vendor/*' -not -path './.git/*' -not -path './cmd' -not -path './.*' -not -path './build/*' -not -path './backup' -not -path './vendor' -not -path '.' -not -path './build' -not -path './etc' -not -path './etc/*' -not -path '*/testdata' -not -path './pkg' | sed 's/^\.\///g'`; do go vet github.com/$(PKG)/$$i; done

test:
	@for i in `find . -type f -name '*_test.go' -not -path './vendor/*' -not -path './.git/*' -not -path '*/testdata'`; do go test -v `dirname $$i`; done
