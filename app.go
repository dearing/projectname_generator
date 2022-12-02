package main

import (
	"fmt"
	"math/rand"
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

func main() {
	fmt.Println(GenerateProjectName())
}
