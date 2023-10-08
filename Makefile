BINDIR:=bin

ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
COMMAND_PACKAGES = $(shell ls -d `pwd`/cmd/*)

# output binary file paths (bin/grpc, bin/rest and so on.)
BINARIES:=$(COMMAND_PACKAGES:$(ROOT_DIR)/cmd/%=$(BINDIR)/%)

# go files list
GO_FILES:=$(shell find . -type f -name '*.go' -print)
# proto files list
PB_FILES:=$(shell find . -type f -name '*.proto' -print)
# generated .go files from proto
GOPB_FILES:=$(PB_FILES:%.proto=%.pb.go)

# build
.PHONY: build
build: $(BINARIES)

# build tasks
$(BINARIES): $(GO_FILES) $(GOPB_FILES)
	@go build -o $@ $(@:$(BINDIR)/%=$(ROOT_DIR)/cmd/%)

# build proto
$(GOPB_FILES): $(PB_FILES) $(BINDIR)/protoc-gen-go
	@protoc \
		--plugin=protoc-gen-go=$(BINDIR)/protoc-gen-go \
		-I ./proto \
		--go_out=./proto \
		--go_opt=paths=source_relative \
		$(@:%.pb.go=%.proto)

$(BINDIR)/protoc-gen-go: go.sum
	@go build -o $@ google.golang.org/protobuf/cmd/protoc-gen-go
