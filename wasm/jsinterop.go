// +build js,wasm
package main

import (
	"syscall/js"
)

func main() {
	js.Global().Get("console").Call("log", "Hello world Go/wasm!")
	js.Global().Get("document").Call("getElementById", "app").Set("innerText", "Hello world")
	js.Global().Get("alert").Invoke("Hello world")
}
