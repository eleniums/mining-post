#!/usr/bin/env bash
set -e

BASE_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )/.." && pwd )"

if [ -f "${BASE_DIR}/scripts/local.test.env" ]; then
  set -o allexport
  source "${BASE_DIR}/scripts/local.test.env"
  set +o allexport
fi

cd $BASE_DIR

go test -v -count=1 -short ./tests $@