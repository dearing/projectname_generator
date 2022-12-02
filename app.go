package main

import (
  "fmt"
  "math/rand"
  "time"
  "syscall/js"
)

var (
  adjectives = []string{
    "gigantic",
    "funky",
    "gleaming",
    "magnificent",
    "majestic",
  }

  nouns = []string{
    "tiger",
    "elephant",
    "unicorn",
    "dragon",
    "pegasus",
  }
)

func main() {
  rand.Seed(time.Now().UnixNano())

  c := make(chan struct{})
  js.Global().Set("generateProjectName", js.FuncOf(generateProjectName))
  <-c
}

func generateProjectName(this js.Value, args []js.Value) interface{} {
  adjective := adjectives[rand.Intn(len(adjectives))]
  noun := nouns[rand.Intn(len(nouns))]
  projectName := fmt.Sprintf("%s %s", adjective, noun)

  return projectName
}
