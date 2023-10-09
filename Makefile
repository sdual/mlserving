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
# generated .go grpc files from proto
GO_GRPCPB_FILES:=$(PB_FILES:%.proto=%_grpc.pb.go)

# GOPATH
GOPATH:=$(shell go env GOPATH)

# build
.PHONY: build
build: $(BINARIES)

.PHONY: clean
clean:
	@$(RM) $(GOPB_FILES) $(GO_GRPCPB_FILES) $(BINARIES)

# build tasks
$(BINARIES): $(GO_FILES) $(GOPB_FILES)
	@GOOS=linux GOARCH=arm64 go  build -o $@ $(@:$(BINDIR)/%=$(ROOT_DIR)/cmd/%)

# build proto
$(GOPB_FILES): $(PB_FILES)
	@protoc \
		--plugin=protoc-gen-go=$(GOPATH)/bin/protoc-gen-go \
		-I ./proto \
		--go_out=./proto \
		--go_opt=paths=source_relative \
		--go-grpc_out=./proto \
		--go-grpc_opt=paths=source_relative \
		$(PB_FILES)
