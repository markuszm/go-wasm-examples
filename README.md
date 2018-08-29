# Go Wasm Examples for Gophers Darmstadt Talk

## Talk

https://speakerdeck.com/markuszm/gophers-darmstadt-running-go-in-your-browser-using-wasm

## Server

Run server to serve wasm binary

`go run ./server/server.go`

## Build examples

Build wasm examples (each file is separate example) with:

`GOOS=js GOARCH=wasm go build -o ./out/test.wasm <example file>`
