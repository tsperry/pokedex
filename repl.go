package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {

	text = strings.ToLower(text)
	return strings.Fields(text)
}

func startRepl(config *config) error {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}
		commandName := input[0]
		command := commands[commandName]
		err := command.callback(config)
		if err != nil {
			fmt.Println("Unknown command")
		}

	}
}
