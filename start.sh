#! /usr/bin/env bash

pushd ./h2c_app
  go build
popd

./h2c_app/h2c-app &

# Hint: you need envoy
envoy --config-path ./envoy/envoy.yaml -l debug
