package main

import (
	"syscall/js"
)

func add(this js.Value, params []js.Value) interface{} {
	sum := 0
	for _, param := range params {
		sum += param.Int()
	}
	return js.ValueOf(sum)
}

func main() {
	c := make(chan struct{}, 0)
	js.Global().Set("add", js.FuncOf(add))
	<-c
}
