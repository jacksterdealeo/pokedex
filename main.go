package main

import (
	"bufio"
	"fmt"
	"os"
)

var page = config{Next: "https://pokeapi.co/api/v2/location-area/"}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		cleanedInput := cleanInput(scanner.Text())
		if len(cleanedInput) < 1 {
			continue
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
			break
		}
		command := cliCommands()[cleanedInput[0]].callback
		if command == nil {
			fmt.Println("Unknown command")
			continue
		} else {
			err := command(&page)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
		// fmt.Println("Your command was:", cleanedInput[0])
	}
}
