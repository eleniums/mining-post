#!/usr/bin/env bash
set -e

BASE_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )/.." && pwd )"

cd $BASE_DIR

NAME=miningpost
VERSION=$(git describe --tags --always --long --dirty)
WINDOWS=${NAME}_windows_amd64.exe
LINUX=${NAME}_linux_amd64
DARWIN=${NAME}_darwin_amd64

echo "Build with version: $VERSION"

echo "Running unit tests..."
./scripts/test_unit.sh

echo "Building for Windows..."
env GOOS=windows GOARCH=amd64 go build -o $WINDOWS -ldflags="-s -w -X main.version=$VERSION" ./cmd/main.go
if [[ -f "$WINDOWS" ]]; then
    echo "Windows build was successful: $WINDOWS"
fi

echo "Building for Linux..."
env GOOS=linux GOARCH=amd64 go build -o $LINUX -ldflags="-s -w -X main.version=$VERSION" ./cmd/main.go
if [[ -f "$LINUX" ]]; then
    echo "Linux build was successful: $LINUX"
fi

echo "Building for macOS..."
env GOOS=darwin GOARCH=amd64 go build -o $DARWIN -ldflags="-s -w -X main.version=$VERSION" ./cmd/main.go
if [[ -f "$DARWIN" ]]; then
    echo "macOS build was successful: $DARWIN"
fi
