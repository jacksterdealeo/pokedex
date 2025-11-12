package main

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"os"

	"github.com/jacksterdealeo/pokedex/internal/api"
)

type cliCommand struct {
	name        string
	description string
	callback    func(config *Config) error
}

func cliCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"catch": {
			name:        "catch",
			description: "TODO: Attempts to catch a Pokemon",
			callback:    commandCatch,
		},
		"exit": {
			name:        "exit",
			description: "Exits the Pokedex",
			callback:    commandExit,
		},
		"explore": {
			name:        "explore",
			description: "Shows all the pokemon in an area",
			callback:    commandExplore,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"inspect": {
			name:        "inspect",
			description: "Displays the name, height, weight, stats and type(s) of the Pokemon.",
			callback:    commandInspect,
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

func commandCatch(config *Config) error {
	var body []byte
	var err error
	if len(config.Parameter) == 0 {
		return fmt.Errorf("no parameter for catching")
	}
	endpoint := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%v/", config.Parameter)

	body, err = api.GetAPIResponse(endpoint, config.Cache)
	if err != nil {
		return err
	}

	var response api.PokemonResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return fmt.Errorf("Couldn't Unmarshal json body\nerr: %v\njson: %v", err, body)
	}

	caughtChance := rand.IntN(response.BaseExperience+250) - response.BaseExperience
	caught := caughtChance > 0

	fmt.Printf("Throwing a Pokeball at %v...\n", config.Parameter)
	//fmt.Println("Debug ~ Rand Chance:", caughtChance, "exp:", response.BaseExperience)
	if caught {
		fmt.Println(config.Parameter, "was caught!")
		config.CaughtPokemon[config.Parameter] = response
	} else {
		fmt.Println(config.Parameter, "escaped!")
	}

	return nil
}

func commandExit(_ *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandExplore(config *Config) error {
	var body []byte
	var err error
	if len(config.Parameter) == 0 {
		return fmt.Errorf("no parameter for exploration")
	}
	endpoint := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%v/", config.Parameter)

	body, err = api.GetAPIResponse(endpoint, config.Cache)
	if err != nil {
		return err
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
	body, err = api.GetAPIResponse(config.Next, config.Cache)
	if err != nil {
		return err
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

func commandInspect(config *Config) error {
	pokemon, exists := config.CaughtPokemon[config.Parameter]
	if !exists {
		return fmt.Errorf("you have not caught that pokemon\n")
	}
	fmt.Printf(`Name: %v
Height: %v
Weight: %v
Stats:
`, pokemon.Name, pokemon.Height, pokemon.Weight)
	for _, e := range pokemon.Stats {
		fmt.Printf("  - %v: %v\n", e.Stat.Name, e.BaseStat)
	}
	fmt.Println("Types:")
	for _, e := range pokemon.Types {
		fmt.Printf("  - %v\n", e.Type.Name)
	}
	return nil
}

func commandMapBack(config *Config) error {
	if config.Previous == "" {
		return fmt.Errorf("There is no previous map")
	}

	var body []byte
	var err error
	body, err = api.GetAPIResponse(config.Previous, config.Cache)
	if err != nil {
		return err
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
