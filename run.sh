#!/usr/bin/bash

set -e

GOOS=js GOARCH=wasm go build -o resources/wasm/bond.wasm ./cmd/main.go
goexec 'http.ListenAndServe(`:8080`, http.FileServer(http.Dir(`./resources/wasm`)))'
