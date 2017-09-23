#!/usr/bin/env bash
mkdir -p bin
CGO_ENABLED=0 GOOS=linux go build -o bin/todolist src/todolist/todolist.go
docker build -t to-do-list .

