#!/bin/bash
SH_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
REPO_DIR=$(dirname $SH_DIR)

go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

golangci-lint run $REPO_DIR/...
