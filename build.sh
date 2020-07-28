#!/bin/bash

ls -lat

go env

GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -o family_finance main.go

if [ ! $? -eq 0 ]; then
  ehco "Go build error"
fi