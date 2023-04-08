#!/usr/bin/env bash

# build if chosen
if [[ "$1" == -c || "$1" == --compile ]]; then
  shift
  make build
fi
# build if necessary
[[ -f ./bin/time-api ]] || make build

# run
./bin/time-api "$@"