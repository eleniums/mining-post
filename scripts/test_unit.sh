#!/usr/bin/env bash
set -e

go test -count=1 -cover ./... $@