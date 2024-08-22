#!/bin/bash
os="$1"
platform="$2"
echo "$os"
echo "$platform"
if [ "$os" = 'linux' ]; then
    GOOS=linux GOARCH=amd64 go build -o ./bin/redisCmd ./main.go
fi

if [ "$os" = 'macos' ]; then
    GOOS=darwin GOARCH=amd64 go build -o ./bin/redisCmd ./main.go
fi

if [ "$os" = 'macos' ] && [ "$platform" = 'arm' ]; then
    GOOS=darwin GOARCH=arm64 go build -o ./bin/redisCmd ./main.go
fi

if [ "$os" = 'windows' ]; then
    GOOS=windows GOARCH=amd64 go build -o ./bin/redisCmd.exe ./main.go
fi
