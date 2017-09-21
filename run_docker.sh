#!/usr/bin/env bash

docker rm -f to-do-list 2> /dev/null || true
docker run --name to-do-list -P to-do-list

