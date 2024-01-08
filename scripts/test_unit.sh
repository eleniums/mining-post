#!/usr/bin/env bash
set -e

go test $(go list ./... | grep -v tests) -count=1 -cover $@