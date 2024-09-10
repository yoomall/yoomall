#!/bin/bash

GOOS=linux GOARCH=amd64 go build -o ./dist/ ./cmd/server
GOOS=linux GOARCH=amd64 go build -o ./dist/ ./cmd/seed