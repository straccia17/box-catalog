#!/bin/sh

echo "Build containers"

if [[ $1 = "PROD" ]]
then
    docker-compose build
else
    docker-compose -f compose.dev.yml build
fi
