package main

import (
	"time"

	"github.com/jacksterdealeo/pokedex/internal/cache"
)

type Config struct {
	Command   string
	Parameter string

	Next     string // URL
	Current  string // URL
	Previous string // URL
	Cache    *cache.Cache
}

func NewConfig() (config *Config) {
	c := Config{
		Next:  "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
		Cache: cache.NewCache(time.Duration(time.Second) * 30),
	}

	return &c
}
