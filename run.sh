#!/usr/bin/bash

set -e

cp $GOROOT/misc/wasm/wasm_exec.js resources/wasm/
GOOS=js GOARCH=wasm go build -o resources/wasm/bond.wasm ./cmd/main.go
goexec 'http.ListenAndServe(`:8080`, http.FileServer(http.Dir(`./resources/wasm`)))'
