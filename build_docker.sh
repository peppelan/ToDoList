#!/usr/bin/env bash

CGO_ENABLED=0 GOOS=linux go build src/todolist/todolist.go
docker build -t to-do-list .

