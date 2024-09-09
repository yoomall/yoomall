FROM golang:1.23.1-bullseye

WORKDIR /app

COPY . .

RUN go mod download

RUN ./build.sh