#!/usr/bin/env bash
set -e
rm -rf bin && mkdir -p bin
CGO_ENABLED=0 GOOS=linux go build -o bin/todolist todolist
docker build -t to-do-list .

