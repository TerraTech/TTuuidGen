PROG          := $(shell basename ${PWD})
VERSION       := $(shell git describe --tags --always HEAD)
VERSION_SHORT := $(shell git describe --abbrev=0 --tags --always HEAD)
BUILD         := $(shell date '+%Y%m%d@%T')
LIB           :=
PKG           := github.com/TerraTech/$(PROG)

CLEAN := $(PROG) version_autogen.go

D_VENDOR	:= $(PWD)/vendor
F_AUTOGEN_LIB	:=
F_AUTOGEN_MAIN	:= ./version_autogen.go
PKG_FQVERSION	:= github.com/TerraTech/FQversion
P_GENVERSION	:= $(D_VENDOR)/$(PKG_FQVERSION)/tools/genVersion.go
P_MAKE_TTUUID	:= make --no-print-directory -C pkg/TTuuid

GOFILES := *.go pkg/*/*.go

all: small

$(PROG): vgen $(GOFILES)
	@go build

.PHONY: build
build: $(PROG)

.PHONY: clean
clean:
	@rm $(CLEAN) >/dev/null 2>&1 || true

.PHONY: fmt
fmt:
	@go fmt ./...
	@$(P_MAKE_TTUUID) fmt

.PHONY: small
small: vgen $(GOFILES)
	@go build -ldflags="-s -w"

.PHONY: test
test:
	@$(P_MAKE_TTUUID) test

.PHONY: vgen
vgen:
	@test -f $(F_AUTOGEN_MAIN) && /bin/rm $(F_AUTOGEN_MAIN) 2>/dev/null || true
	@PROG=$(PROG) VERSION=$(VERSION) BUILD=$(BUILD) IMPFQVERSION=$(PKG_FQVERSION) \
	  go generate $(dir $(F_AUTOGEN_MAIN))version.go
