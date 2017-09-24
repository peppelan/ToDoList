#!/usr/bin/env bash

docker rm -f to-do-list 2> /dev/null || true
docker run -d --name to-do-list -p8080:8080 to-do-list

