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
		"Aardvark",
		"Albatross",
		"Ant",
		"Baboon",
		"Badger",
		"Beaver",
		"Cheetah",
		"Cougar",
		"Crocodile",
		"Dingo",
		"Dolphin",
		"Dove",
		"Dragon",
		"Eagle",
		"Elephant",
		"Emu",
		"Flamingo",
		"Fossa",
		"Fox",
		"Gazelle",
		"Giraffe",
		"Heron",
		"Hippopotamus",
		"Hyena",
		"Ibis",
		"Iguana",
		"Impala",
		"Jackal",
		"Jaguar",
		"Jellyfish",
		"Kangaroo",
		"Kiwi",
		"Lemur",
		"Leopard",
		"Lion",
		"Meerkat",
		"Mongoose",
		"Monkey",
		"Newt",
		"Nighthawk",
		"Numbat",
		"Ocelot",
		"Octopus",
		"Opossum",
		"Panda",
		"Panther",
		"Puma",
		"Quail",
		"Quokka",
		"Quoll",
		"Rabbit",
		"Raccoon",
		"Seahorse",
		"Sloth",
		"Squirrel",
		"Tarantula",
		"Tiger",
		"Tortoise",
		"Uakari",
		"Urutu",
		"Viper",
		"Vole",
		"Vole",
		"Whale",
		"Wombat",
		"Wombat",
		"Xerus",
		"Yak",
		"Yak",
		"Zebra",
		"Zorilla",
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
