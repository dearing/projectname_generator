//go:build js && wasm

package main

import (
	"fmt"
	"math/rand"
	"strings"
	"syscall/js"
	"time"
)

var (
	adjectives = []string{
		"adorable",
		"alluring",
		"amazing",
		"awesome",
		"beautiful",
		"breathtaking",
		"brilliant",
		"classy",
		"clever",
		"cool",
		"dazzling",
		"delightful",
		"elegant",
		"enchanting",
		"exquisite",
		"fabulous",
		"fantastic",
		"gorgeous",
		"graceful",
		"grand",
		"great",
		"handsome",
		"incredible",
		"magnificent",
		"marvelous",
		"mesmerizing",
		"majestic",
		"outstanding",
		"perfect",
		"pleasant",
		"pleasing",
		"powerful",
		"radiant",
		"refined",
		"remarkable",
		"resplendent",
		"sensational",
		"splendid",
		"stunning",
		"sublime",
		"superb",
		"terrific",
		"transcendent",
		"wonderful",
		"xenodochial",
		"yummy",
		"zesty",
	}
	nouns = []string{
		"aardvark",
		"albatross",
		"ant",
		"baboon",
		"badger",
		"beaver",
		"cheetah",
		"cougar",
		"crocodile",
		"dingo",
		"dolphin",
		"dove",
		"dragon",
		"eagle",
		"elephant",
		"emu",
		"flamingo",
		"fossa",
		"fox",
		"gazelle",
		"giraffe",
		"heron",
		"hippopotamus",
		"hyena",
		"ibis",
		"iguana",
		"impala",
		"jackal",
		"jaguar",
		"jellyfish",
		"kangaroo",
		"kiwi",
		"lemur",
		"leopard",
		"lion",
		"meerkat",
		"mongoose",
		"monkey",
		"newt",
		"nighthawk",
		"numbat",
		"ocelot",
		"octopus",
		"opossum",
		"panda",
		"panther",
		"puma",
		"quail",
		"quokka",
		"quoll",
		"rabbit",
		"raccoon",
		"seahorse",
		"sloth",
		"squirrel",
		"tarantula",
		"tiger",
		"tortoise",
		"uakari",
		"urutu",
		"viper",
		"vole",
		"vole",
		"whale",
		"wombat",
		"wombat",
		"xerus",
		"yak",
		"yak",
		"zebra",
		"zorilla",
	}
)

// GenerateProjectName generates a random project name.
func GenerateProjectName() string {
	rand.Seed(time.Now().Unix())
	return strings.ToLower(fmt.Sprintf("%s_%s", adjectives[rand.Intn(len(adjectives))], nouns[rand.Intn(len(nouns))]))
}

func gernerateFunction(this js.Value, p []js.Value) interface{} {
	GenerateProjectName()
	return js.ValueOf(GenerateProjectName())
}

func main() {
	println("Hello, WebAssembly!")
	c := make(chan struct{}, 0)
	for i := 0; i < 10; i++ {
		fmt.Println(GenerateProjectName())
	}
	js.Global().Set("generate", js.FuncOf(gernerateFunction))
	<-c
}
