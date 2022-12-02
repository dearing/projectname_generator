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

func GenerateProjectName() string {
	return strings.ToLower(fmt.Sprintf("%s_%s", adjectives[rand.Intn(len(adjectives))], nouns[rand.Intn(len(nouns))]))
}

func generateFunction(this js.Value, args []js.Value) interface{} {
	count := 1
	if len(args) != 0 {
		count = args[0].Int()
	}
	var names []interface{}
	for i := 0; i < count; i++ {
		names = append(names, GenerateProjectName())
	}

	return js.ValueOf(names[:])
}

func main() {
	println("app loaded")
	rand.Seed(time.Now().Unix())
	c := make(chan struct{}, 0)
	js.Global().Set("generate", js.FuncOf(generateFunction))
	<-c
}
