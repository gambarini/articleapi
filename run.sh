#!/usr/bin/env bash

go build ./...

echo "starting mongodb..."
docker run -p 27017:27017 --name mongodb_articleapi -d mongo:4.0.0

echo "waiting mongodb..."
sleep 5s

go run main.go

echo "stoping mongodb"
docker stop mongodb_articleapi

echo "removing containers"
docker rm mongodb_articleapi
