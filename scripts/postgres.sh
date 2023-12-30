#!/bin/bash
set -e

CONTAINER_NAME=mining-post-db
DB_PORT=5432

BASE_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )/.." && pwd )"

# need a file postgres.env with DB_PASSWORD inside
set -o allexport
source "${BASE_DIR}/scripts/postgres.env"
set +o allexport

# run postgres container on default port if not already running
if !(echo >/dev/tcp/127.0.0.1/$DB_PORT) &>/dev/null
then
    # try starting container first, then create fresh if it does not exist
    # first time running this will have errors from the first command, but it is still successful
    docker start $CONTAINER_NAME || docker run -p $DB_PORT:$DB_PORT --name $CONTAINER_NAME -e POSTGRES_PASSWORD=${DB_PASSWORD} -d postgres
fi

# display container status
docker ps -a | grep $CONTAINER_NAME