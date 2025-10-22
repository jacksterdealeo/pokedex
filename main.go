package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

/*
The purpose of this function will be to split the user's input into "words" based on whitespace.
It should also lowercase the input and trim any leading or trailing whitespace.
For example:

	hello world -> ["hello", "world"]
	Charmander Bulbasaur PIKACHU -> ["charmander", "bulbasaur", "pikachu"]
*/
func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
