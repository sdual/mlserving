BINDIR:=bin

ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
COMMAND_PACKAGES = $(shell ls -d `pwd`/cmd/*)

# output binary file paths (bin/grpc, bin/rest and so on.)
BINARIES:=$(COMMAND_PACKAGES:$(ROOT_DIR)/cmd/%=$(BINDIR)/%)

# go files list
GO_FILES:=$(shell find . -type f -name '*.go' -print)

# build
.PHONY: build
build: $(BINARIES)

# build tasks
$(BINARIES): $(GO_FILES)
	@go build -o $@ $(@:$(BINDIR)/%=$(ROOT_DIR)/cmd/%)

