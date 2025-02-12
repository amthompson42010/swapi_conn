package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
)

func main() {
	// Define the API endpoint URL
	url := "https://swapi.dev/api/people/1"

	// Create a new HTTP client
	client := &http.Client{}

	// Create a new GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

    // Optionally add headers to the request
	req.Header.Add("Accept", "application/json")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Print the response body
	fmt.Println(string(body))
}