SHELL := /bin/bash

PROG          := TTuuidGen
VERSION       := $(shell git describe --tags --always HEAD)
VERSION_SHORT := $(shell git describe --abbrev=0 --tags --always HEAD)
BUILD         := $(shell date '+%Y%m%d@%T')
LIB           :=
PKG           := github.com/TerraTech/$(PROG)

GIT_COMMIT = $(shell git rev-parse HEAD)
GIT_SHA    = $(shell git rev-parse --short HEAD)
GIT_TAG    = $(shell git describe --tags --abbrev=0 --exact-match 2>/dev/null)
GIT_DIRTY  = $(shell [[ -n "$(git status --porcelain)" ]] && echo "dirty" || echo "clean")

D_BIN          := $(CURDIR)/bin
D_CMD          := ./cmd/$(PROG)
D_PKG          := ./pkg/TTuuid
F_AUTOGEN_MAIN := $(D_CMD)/version_autogen.go
PKG_FQVERSION  := github.com/TerraTech/FQversion
P_MAKE_TTUUID  := make --no-print-directory -C $(D_PKG)

# go option
PKG     := ./...
LDFLAGS :=
GOFLAGS :=

CLEAN   := $(D_BIN)/$(PROG) $(F_AUTOGEN_MAIN)
GOFILES := $(D_CMD)/*.go $(D_PKG)/*.go

all: small

$(D_BIN)/$(PROG): vgen $(GOFILES)
	@GO111MODULE=on go build $(GOFLAGS) -ldflags '$(LDFLAGS)' -o $(D_BIN)/$(PROG) $(D_CMD)

.PHONY: build
build: $(D_BIN)/$(PROG)

.PHONY: clean
clean:
	@rm $(CLEAN) >/dev/null 2>&1 || true

.PHONY: fmt
fmt:
	@go fmt ./...
	@$(P_MAKE_TTUUID) fmt

.PHONY: install
install:
	@go install $(D_CMD)

.PHONY: small
small: vgen $(GOFILES)
	@LDFLAGS="-s -w" make -s build

.PHONY: tag
tag:
ifndef TAG
	$(error TAG is undefined)
else
	@git rtag "$(TAG)"
	@while read m; do git tag "$${m%/go.mod}/$(TAG)"; done < <(shopt -s globstar; ls -1 pkg/**/go.mod)
endif

.PHONY: test
test:
	@$(P_MAKE_TTUUID) test

.PHONY: vgen
vgen:
	@test -f $(F_AUTOGEN_MAIN) && /bin/rm $(F_AUTOGEN_MAIN) 2>/dev/null || true
	@PROG=$(PROG) VERSION=$(VERSION) BUILD=$(BUILD) IMPFQVERSION=$(PKG_FQVERSION) \
	  go generate $(dir $(F_AUTOGEN_MAIN))version.go
