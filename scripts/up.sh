#!/bin/sh

echo "Start containers"

if [[ $1 = "PROD" ]]
then
    docker-compose up
else
    docker-compose -f compose.dev.yml up
fi
