package main

import "fmt"

func commandHelp(config *config) error {
	message := `
Welcome to the Pokedex!
Usage:

`
	for _, command := range config.commands {
		message = message + command.name + ": " + command.description + "\n"
	}
	fmt.Println(message)
	return nil
}
