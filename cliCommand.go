package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/jacksterdealeo/pokedex/internal/api"
	"github.com/jacksterdealeo/pokedex/internal/pokecache"
)

var cache = pokecache.NewCache(time.Duration(time.Second) * 30)

type cliCommand struct {
	name        string
	description string
	callback    func(config *Config) error
}

func cliCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exits the Pokedex",
			callback:    commandExit,
		},
		"explore": {
			// TODO: MAKE THIS
			name:        "explore",
			description: "TODO: Shows all the pokemon in an area",
			callback:    commandExplore,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Shows the next area",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Shows the previous area",
			callback:    commandMapBack,
		},
		"quit": {
			name:        "quit",
			description: "Does the same as 'exit'",
			callback:    commandExit,
		},
	}
}

func commandExit(_ *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandExplore(config *Config) error {
	// TODO: MAKE THIS
	var body []byte
	var err error
	if len(config.Parameter) == 0 {
		return fmt.Errorf("no parameter for exploration")
	}
	endpoint := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%v/", config.Parameter)
	if data, found := cache.Get(endpoint); found {
		body = data
	} else {
		body, err = api.GetAPIResponse(endpoint)
		if err != nil {
			return err
		}
		cache.Add(endpoint, body)
	}

	var response api.ExploreResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return fmt.Errorf("Couldn't Unmarshal json body\nerr: %v\njson: %v", err, body)
	}

	for _, result := range response.PokemonEncounters {
		fmt.Println(result.Pokemon.Name)
	}

	return nil
}

func commandHelp(_ *Config) error {
	fmt.Println(`Welcome to the Pokedex!
Usage:`)
	for _, value := range cliCommands() {
		fmt.Printf("\t%v: %v\n", value.name, value.description)
	}
	return nil
}

func commandMap(config *Config) error {
	var body []byte
	var err error
	if data, found := cache.Get(config.Next); found {
		body = data
	} else {
		body, err = api.GetAPIResponse(config.Next)
		if err != nil {
			return err
		}
		cache.Add(config.Next, body)
	}

	var response api.MapResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return fmt.Errorf("Couldn't Unmarshal json body\nerr: %v\njson: %v", err, body)
	}

	for _, result := range response.Results {
		fmt.Println(result.Name)
	}

	config.Previous = response.Previous
	config.Current = config.Next
	config.Next = response.Next
	return nil
}

func commandMapBack(config *Config) error {
	if config.Previous == "" {
		return fmt.Errorf("There is no previous map")
	}

	var body []byte
	var err error
	if data, found := cache.Get(config.Previous); found {
		body = data
	} else {
		body, err = api.GetAPIResponse(config.Previous)
		if err != nil {
			return err
		}
		cache.Add(config.Previous, body)
	}

	var response api.MapResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return fmt.Errorf("Couldn't Unmarshal json body, %v", err)
	}

	for _, result := range response.Results {
		fmt.Println(result.Name)
	}

	config.Current = config.Previous
	config.Previous = response.Previous
	config.Next = response.Next
	return nil
}
