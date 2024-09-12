#!/bin/bash

# linux
GOOS=linux GOARCH=amd64 go build -o ./dist/linux/ ./cmd/server
GOOS=linux GOARCH=amd64 go build -o ./dist/linux/ ./cmd/seed

# windows 
GOOS=windows GOARCH=amd64 go build -o ./dist/win/ ./cmd/server
GOOS=windows GOARCH=amd64 go build -o ./dist/win/ ./cmd/seed

# macos 
GOOS=darwin GOARCH=amd64 go build -o ./dist/mac/ ./cmd/server
GOOS=darwin GOARCH=amd64 go build -o ./dist/mac/ ./cmd/seed


# move dev.config.yaml to dist/{plat}/config.yaml
cp dev.config.yaml ./dist/linux/config.yaml
cp dev.config.yaml ./dist/win/config.yaml
cp dev.config.yaml ./dist/mac/config.yaml