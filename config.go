package main

import (
	"time"

	"github.com/jacksterdealeo/pokedex/internal/pokecache"
)

type Config struct {
	Command   string
	Parameter string

	Next     string // URL
	Current  string // URL
	Previous string // URL
	Cache    *pokecache.Cache
}

func NewConfig() (config *Config) {
	c := Config{
		Next:  "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
		Cache: pokecache.NewCache(time.Duration(time.Second) * 30),
	}

	return &c
}
