linux:
	GOOS=linux GOARCH=amd64 go build -o ./dist/linux/ ./cmd/server
	GOOS=linux GOARCH=amd64 go build -o ./dist/linux/ ./cmd/tools
	GOOS=linux GOARCH=amd64 go build -o ./dist/linux/ ./cmd/jobs
	cp dev.config.yaml ./dist/linux/config.yaml

windows:
	GOOS=windows GOARCH=amd64 go build -o ./dist/windows/ ./cmd/server
	GOOS=windows GOARCH=amd64 go build -o ./dist/windows/ ./cmd/tools
	GOOS=windows GOARCH=amd64 go build -o ./dist/windows/ ./cmd/jobs
	cp dev.config.yaml ./dist/windows/config.yaml

macos:
	GOOS=darwin GOARCH=amd64 go build -o ./dist/mac/ ./cmd/server
	GOOS=darwin GOARCH=amd64 go build -o ./dist/mac/ ./cmd/tools
	GOOS=darwin GOARCH=amd64 go build -o ./dist/mac/ ./cmd/jobs
	cp dev.config.yaml ./dist/mac/config.yaml

auth_service:
	GOOS=linux GOARCH=amd64 go build -o ./dist/linux/ ./cmd/service/auth
	cp dev.config.yaml ./dist/linux/config.yaml

run:
	go run ./cmd/server

jobs:
	go run ./cmd/jobs

wire:
	wire ./cmd/server
	wire ./cmd/tools
	wire ./cmd/service/auth
	wire ./cmd/jobs

test:
	go test -v ./...


prod:
	make wire
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./dist/linux/ ./cmd/server && upx -9 ./dist/linux/server
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./dist/linux/ ./cmd/tools && upx -9 ./dist/linux/tools
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./dist/linux/ ./cmd/jobs && upx -9 ./dist/linux/jobs
	cp -r ./templates ./dist/linux/templates
	cp -r ./public ./dist/linux/public