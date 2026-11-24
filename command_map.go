package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Result struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Results  []Location `json:"results"`
}

type Location struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func commandMap(config *config) error {
	getUrl := "https://pokeapi.co/api/v2/location-area/"
	if config.next != "" {
		getUrl = config.next
	}

	res, err := http.Get(getUrl)
	if err != nil {
		log.Fatal(err)
		return err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return err
	}

	results := Result{}
	err = json.Unmarshal(body, &results)
	if err != nil {
		log.Fatal(err)
		return err
	}
	// fmt.Printf("results.count: %v\n", results.Count)
	// fmt.Printf("results.next: %v\n", results.Next)
	// fmt.Printf("results.previous: %v\n", results.Previous)
	// fmt.Printf("results.results: %v\n", results.Results)
	config.next = results.Next
	config.previous = results.Previous
	locations := results.Results
	for _, location := range locations {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapBack(config *config) error {
	if config.previous == "" {
		fmt.Println("You're at the first page")
		return nil
	}
	getUrl := config.previous

	res, err := http.Get(getUrl)
	if err != nil {
		log.Fatal(err)
		return err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return err
	}

	results := Result{}
	err = json.Unmarshal(body, &results)
	if err != nil {
		log.Fatal(err)
		return err
	}
	// fmt.Printf("results.count: %v\n", results.Count)
	// fmt.Printf("results.next: %v\n", results.Next)
	// fmt.Printf("results.previous: %v\n", results.Previous)
	// fmt.Printf("results.results: %v\n", results.Results)
	config.next = results.Next
	config.previous = results.Previous
	locations := results.Results
	for _, location := range locations {
		fmt.Println(location.Name)
	}
	return nil
}
