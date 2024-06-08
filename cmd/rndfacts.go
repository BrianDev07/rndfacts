package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"rndfacts/api"
	"time"
)

const (
	url      = "https://catfact.ninja/fact"
	filepath = "./facts.json"
)

func main() {
	client := &http.Client{Timeout: 10 * time.Second}
	facts := new(api.Facts)

	jsonBytes, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	if !json.Valid(jsonBytes) {
		panic(fmt.Errorf("file '%v' is not a valid json", filepath))
	}

	if err := json.Unmarshal(jsonBytes, &facts); err != nil {
		panic(err)
	}

	randomCatFact, err := api.GetCatFact(client, url)
	if err != nil {
		panic(err)
	}

	facts.Facts = append(facts.Facts, randomCatFact)

	jsonFacts, err := json.MarshalIndent(facts, "", "    ")
	if err != nil {
		panic(err)
	}

	if writeError := os.WriteFile(filepath, jsonFacts, 0644); writeError != nil {
		panic(writeError)
	}

	fmt.Printf("Fact: %v\n", facts.Facts[len(facts.Facts)-1].Fact)
}
