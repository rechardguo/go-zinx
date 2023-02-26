#!/bin/bash
GOOS=windows
GOARCH=amd64

go build -o server server.go user.go

go build -o client client.go