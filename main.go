package main

import (
	"fmt"
	"syscall/js"
)

func main() {

	c := make(chan int)
	fmt.Println("Go WebAssembly Tutorial")

	document := js.Global().Get("document")

	hello := document.Call("createElement", "h1")
	hello.Set("innetHtml", "Go WebAssembly Demo")
	document.Get("body").Call("appendChild", hello)

	js.Global().Set("add", js.FuncOf(add))

	<-c
}

func add(this js.Value, args []js.Value) interface{} {
	return args[0].Float() + args[1].Float()
}
