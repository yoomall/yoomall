linux:
	GOOS=linux GOARCH=amd64 go build -o ./dist/linux/ ./cmd/server
	GOOS=linux GOARCH=amd64 go build -o ./dist/linux/ ./cmd/seed
	cp dev.config.yaml ./dist/linux/config.yaml

windows:
	GOOS=windows GOARCH=amd64 go build -o ./dist/windows/ ./cmd/server
	GOOS=windows GOARCH=amd64 go build -o ./dist/windows/ ./cmd/seed
	cp dev.config.yaml ./dist/windows/config.yaml

macos:
	GOOS=darwin GOARCH=amd64 go build -o ./dist/mac/ ./cmd/server
	GOOS=darwin GOARCH=amd64 go build -o ./dist/mac/ ./cmd/seed
	cp dev.config.yaml ./dist/mac/config.yaml

auth_service:
	GOOS=linux GOARCH=amd64 go build -o ./dist/linux/ ./cmd/service/auth
	cp dev.config.yaml ./dist/linux/config.yaml

run:
	go run ./cmd/server

wire:
	wire ./cmd/server
	wire ./cmd/seed
	wire ./cmd/service/auth

test:
	go test -v ./...


prod_linux:
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./dist/linux/ ./cmd/server && upx -9 ./dist/linux/server
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./dist/linux/ ./cmd/seed && upx -9 ./dist/linux/seed