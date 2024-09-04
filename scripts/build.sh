#!/bin/sh

set -e
set -x

# Ensure that Go modules are used
export GOPATH=""

# Disable cgo
export CGO_ENABLED=0

# Build the Go executable for amd64 architecture
GOOS=linux GOARCH=amd64 go build -o release/linux/amd64/drone-ant-plugin ./main.go
GOOS=linux GOARCH=arm64 go build -o release/linux/arm64/drone-ant-plugin ./main.go

GOOS=windows go build -o release/windows/amd64/drone-ant-plugin.exe