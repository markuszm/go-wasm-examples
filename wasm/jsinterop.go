// +build js,wasm

package main

import (
	"syscall/js"
)

var (
	console  = js.Global().Get("console")
	document = js.Global().Get("document")
	alert    = js.Global().Get("alert")
)

func main() {
	console.Call("log", "Hello world Go/wasm!")

	alert.Invoke("Hello world")

	app := document.Call("getElementById", "app")
	app.Set("innerText", "Hello world")

	document.Call("querySelector", "body").Get("style").Set("background", "green")

}
