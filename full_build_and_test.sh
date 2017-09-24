#!/usr/bin/env bash
set -e

./build_docker.sh
./run_docker.sh
./run_acceptance_tests.sh