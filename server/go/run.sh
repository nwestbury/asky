#!/bin/bash
set -e
export GOPATH=$(pwd)
. bin/install.sh
go build main
./main
