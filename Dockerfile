FROM alpine:latest

WORKDIR /app
COPY ./dist/ /app
ADD ./docker/config.yaml /app/config.yaml
