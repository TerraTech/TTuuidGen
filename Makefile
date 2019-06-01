PROG          := $(shell basename ${PWD})
VERSION       := $(shell git describe --tags --always HEAD)
VERSION_SHORT := $(shell git describe --abbrev=0 --tags --always HEAD)
BUILD         := $(shell date '+%Y%m%d@%T')
LIB           :=
PKG           := github.com/TerraTech/$(PROG)

D_CMD		:= ./cmd/$(PROG)
D_PKG		:= ./pkg/TTuuid
F_AUTOGEN_MAIN	:= $(D_CMD)/version_autogen.go
PKG_FQVERSION	:= github.com/TerraTech/FQversion
P_MAKE_TTUUID	:= make --no-print-directory -C $(D_PKG)

CLEAN := $(PROG) $(F_AUTOGEN_MAIN)
GOFILES := $(D_CMD)/*.go $(D_PKG)/*.go

all: small

$(PROG): vgen $(GOFILES)
	@go build -o $(PROG) $(D_CMD)

.PHONY: build
build: $(PROG)

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
	@go build -ldflags="-s -w" -o $(PROG) $(D_CMD)

.PHONY: test
test:
	@$(P_MAKE_TTUUID) test

.PHONY: vgen
vgen:
	@test -f $(F_AUTOGEN_MAIN) && /bin/rm $(F_AUTOGEN_MAIN) 2>/dev/null || true
	@PROG=$(PROG) VERSION=$(VERSION) BUILD=$(BUILD) IMPFQVERSION=$(PKG_FQVERSION) \
	  go generate $(dir $(F_AUTOGEN_MAIN))version.go
