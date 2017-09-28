#!/usr/bin/env bash
set -e
go test src/acceptance_tests/*.go -args -url=http://localhost
