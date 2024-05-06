#!/bin/bash
SH_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
REPO_DIR=$(dirname $SH_DIR)

go install honnef.co/go/tools/cmd/staticcheck@latest
staticcheck $REPO_DIR/...
