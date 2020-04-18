package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"deadly.surgery/memelord79/internal/state"
)

var repo = "git@github.com:DeadlySurgeon/memelord79.git"

func main() {
	store, err := state.Get(repo)
	if err != nil {
		log.Fatalf("Failed to get state: %v", err)
	}
	defer func() {
		if err = state.Save(repo, store); err != nil {
			log.Fatalf("Failed to save state: %v", err)
		}
	}()

	fact, err := getFact()
	if err != nil {
		log.Fatalf("Failed to get a printer fact: %v\n", err)
	}

	fmt.Println(fact)
	store.Example = fact
}

func getFact() (string, error) {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequest(
		http.MethodGet,
		"https://printerfacts.cetacean.club/fact",
		nil,
	)
	if err != nil {
		return "", err
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	factBytes, err := ioutil.ReadAll(resp.Body)
	return string(factBytes), err
}
