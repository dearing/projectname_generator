//go:build js && wasm

package main

import (
	"fmt"
	"math/rand"
	"syscall/js"
)

var (
	adjectives = []string{
		"adorable",
		"beautiful",
		"clever",
		"elegant",
		"fierce",
		"gorgeous",
		"majestic",
		"powerful",
		"quirky",
		"sparkling",
		"witty",
	}
	nouns = []string{
		"bird",
		"cat",
		"dragon",
		"flower",
		"fox",
		"horse",
		"unicorn",
		"zebra",
	}
)

// GenerateProjectName generates a random project name.
func GenerateProjectName() string {
	return fmt.Sprintf("%s %s", adjectives[rand.Intn(len(adjectives))], nouns[rand.Intn(len(nouns))])
}

func gernerateFunction(this js.Value, p []js.Value) interface{} {
	return js.ValueOf(GenerateProjectName())
}

func main() {
	c := make(chan struct{}, 0)
	//fmt.Println(GenerateProjectName())
	println("Hello, WebAssembly!")
	js.Global().Set("generate", js.FuncOf(gernerateFunction))
	<-c
}
