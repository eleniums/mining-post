#!/usr/bin/env bash
set -e

BASE_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )/.." && pwd )"

if [ -f "${BASE_DIR}/scripts/local.env" ]; then
  set -o allexport
  source "${BASE_DIR}/scripts/local.env"
  set +o allexport
fi

cd $BASE_DIR

go run ./cmd/main.go $@