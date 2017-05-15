#!/bin/bash
set -e
. bin/install.sh
go build main
./main
