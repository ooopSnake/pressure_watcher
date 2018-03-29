#!/usr/bin/env bash
echo "delete previous builds"
rm -rf pressure_watcher_*

buildPrefix="pressure_watcher_"

echo "build for arm linux"
CGO_ENABLED=0 GOOS=linux GOARCH=arm go  build -v -o ${buildPrefix}arm

echo "build for x86 linux"
CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -v -o ${buildPrefix}x86

echo "build for x64 linux"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o ${buildPrefix}x64

echo "all done"