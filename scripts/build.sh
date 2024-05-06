#!/bin/bash
SH_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
REPO_DIR=$(dirname $SH_DIR)

cd $REPO_DIR
go build -o bin/hbk
