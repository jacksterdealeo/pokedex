package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	config := NewConfig()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		cleanedInput := cleanInput(scanner.Text())

		if len(cleanedInput) < 1 {
			continue
		}
		config.Command = cleanedInput[0]
		if len(cleanedInput) >= 2 {
			config.Parameter = cleanedInput[1]
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
			break
		}
		command := cliCommands()[cleanedInput[0]].callback

		if len(cleanedInput) >= 2 {
			config.Parameter = cleanedInput[1]
		} else {
			config.Parameter = ""
		}

		if command == nil {
			fmt.Println("Unknown command")
			continue
		} else {
			err := command(config)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
		// fmt.Println("Your command was:", cleanedInput[0])
	}
}
