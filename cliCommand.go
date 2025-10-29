package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jacksterdealeo/pokedex/internal/api"
	"github.com/jacksterdealeo/pokedex/internal/pokecache"
)

var cache = pokecache.NewCache(time.Duration(time.Second) * 30)

// contains the Next and Previous URLs needed to paginate through location areas.
// When getting Next and Previous values, PokeAPI returns URLs to pages 20 entries long by default.
type config struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type cliCommand struct {
	name        string
	description string
	callback    func(config *config) error
}

func cliCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exits the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "",
			callback:    commandMapBack,
		},
		"quit": {
			name:        "quit",
			description: "An alias for exit",
			callback:    commandExit,
		},
	}
}

func commandExit(_ *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(_ *config) error {
	fmt.Println(`Welcome to the Pokedex!
Usage:`)
	for _, value := range cliCommands() {
		fmt.Printf("\t%v: %v\n", value.name, value.description)
	}
	return nil
}

func commandMap(config *config) error {
	if config.Next == "" {
		return fmt.Errorf("There is no next map")
	}

	var body []byte
	var err error
	if data, found := cache.Get(config.Next); found {
		body = data
	} else {
		body, err = api.GetResponseBody(config.Next)
		if err != nil {
			return err
		}
		cache.Add(config.Next, body)
	}
	err = json.Unmarshal(body, config)
	if err != nil {
		log.Fatalf("Couldn't Unmarshal json body\nerr: %v\njson: %v", err, body)
	}

	for _, result := range config.Results {
		fmt.Println(result.Name)
	}
	return nil
}

func commandMapBack(config *config) error {
	if config.Previous == "" {
		return fmt.Errorf("There is no previous map")
	}

	var body []byte
	var err error
	if data, found := cache.Get(config.Previous); found {
		body = data
	} else {
		body, err = api.GetResponseBody(config.Previous)
		if err != nil {
			return err
		}
		cache.Add(config.Previous, body)
	}
	err = json.Unmarshal(body, config)
	if err != nil {
		log.Fatalf("Couldn't Unmarshal json body, %v", err)
	}

	for _, result := range config.Results {
		fmt.Println(result.Name)
	}
	return nil
}
