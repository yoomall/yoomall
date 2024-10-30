linux:
	GOOS=linux GOARCH=amd64 go build -o ./dist/linux/ ./cmd/http
	GOOS=linux GOARCH=amd64 go build -o ./dist/linux/ ./cmd/tools
	GOOS=linux GOARCH=amd64 go build -o ./dist/linux/ ./cmd/jobs
	cp dev.config.yaml ./dist/linux/config.yaml

windows:
	GOOS=windows GOARCH=amd64 go build -o ./dist/windows/ ./cmd/http
	GOOS=windows GOARCH=amd64 go build -o ./dist/windows/ ./cmd/tools
	GOOS=windows GOARCH=amd64 go build -o ./dist/windows/ ./cmd/jobs
	cp dev.config.yaml ./dist/windows/config.yaml

macos:
	GOOS=darwin GOARCH=amd64 go build -o ./dist/mac/ ./cmd/http
	GOOS=darwin GOARCH=amd64 go build -o ./dist/mac/ ./cmd/tools
	GOOS=darwin GOARCH=amd64 go build -o ./dist/mac/ ./cmd/jobs
	cp dev.config.yaml ./dist/mac/config.yaml

run:
	go run ./cmd/http

jobs:
	go run ./cmd/jobs

wire:
	wire ./cmd/http/api
	wire ./cmd/tools
	wire ./cmd/jobs

test:
	go test -v ./...


prod:
	make wire
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./dist/linux/ ./cmd/http && upx -9 ./dist/linux/http
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./dist/linux/ ./cmd/tools && upx -9 ./dist/linux/tools
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./dist/linux/ ./cmd/jobs && upx -9 ./dist/linux/jobs
	cp -r ./templates ./dist/linux/templates
	cp -r ./public ./dist/linux/public

vercel:
	make wire
	cp -r ./templates ./api/templates
	vercel --prod