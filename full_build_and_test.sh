#!/usr/bin/env bash
set -e

# Unit tests
go test todolist

# Prepare docker container
./build_docker.sh > /dev/null
./run_docker.sh

# Acceptance tests
./run_acceptance_tests.sh
