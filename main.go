package main

import (
	"time"

	"github.com/NHuxoll/pokedex/internal/api"
)

func main() {
	pokeClient := api.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
