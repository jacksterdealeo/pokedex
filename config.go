package main

type Config struct {
	Command   string
	Parameter string

	Next     string // URL
	Current  string // URL
	Previous string // URL
}

func NewConfig() (config *Config) {
	c := Config{
		Next: "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
	}

	return &c
}
