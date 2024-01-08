#!/usr/bin/env bash
set -e

# TODO: fix unit test script to only run unit tests, not integration tests. Use grep
go test -count=1 -cover ./... $@