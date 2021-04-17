#!/usr/bin/env bash
cd ../api
CGO_ENABLE=0 GOOS=linux GOARCH=amd64 go build -o ../build
