#!/usr/bin/env bash
set -e

echo "--- Running Unit tests..."
go test todolist

echo "--- Preparing docker image..."
./build_docker.sh > /dev/null

echo "--- Running docker container..."
./run_docker.sh

echo "--- Running Acceptance tests..."
./run_acceptance_tests.sh
