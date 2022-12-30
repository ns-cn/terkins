#!/bin/zsh
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o ./releases/terkins_linux .
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-w -s" -o ./releases/terkins.exe .
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-w -s" -o ./releases/terkins_darwin .