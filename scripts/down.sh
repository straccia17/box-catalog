#!/bin/sh

echo "Stop and remove containers"

if [[ $1 = "PROD" ]]
then
    docker compose down
else
    docker compose -f compose.dev.yml down
fi
