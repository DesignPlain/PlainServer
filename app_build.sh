#!/usr/bin/env bash

bash app_setup.sh

go mod tidy -v

go build -v -o bin/server ./src
