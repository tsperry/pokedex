package main

import {
	"pokedexcli/internal/pokecache"
}

type config struct {
	commands map[string]cliCommand
	next     string
	previous string
}

type cliCommand struct {
	name        string
	description string
	callback    func(config *config) error
}

var commands = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"help": {
		name:        "help",
		description: "Display a help message",
		callback:    commandHelp,
	},
	"map": {
		name:        "map",
		description: "Display next 20 locations",
		callback:    commandMap,
	},
	"mapb": {
		name:        "mapb",
		description: "Display previous 20 locations",
		callback:    commandMapBack,
	},
}

func main() {
	config := config{
		commands: commands,
	}
	startRepl(&config)
}
