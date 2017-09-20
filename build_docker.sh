#!/usr/bin/env bash

CGO_ENABLED=0 GOOS=linux go build src/hello/hello.go 
docker build -t to-do-list .

