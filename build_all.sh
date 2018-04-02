#!/usr/bin/env bash

echo "delete previous builds"
rm -rf pressure_watcher_*

buildPrefix="pressure_watcher_"

echo "building for arm linux"
CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -ldflags "-w -s" -v -o ${buildPrefix}arm1

echo "building for x86 linux"
CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -ldflags "-w -s" -v -o ${buildPrefix}x86

echo "building for x64 linux"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -v -o ${buildPrefix}x64

echo "all done"