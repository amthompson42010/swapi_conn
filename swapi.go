package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	"github.com/gorilla/mux"
)

type Person struct {
	Name      string   `json:"name"`
	Height    string   `json:"height"`
	Mass      string   `json:"mass"`
	HairColor string   `json:"hair_color"`
	SkinColor string   `json:"skin_color"`
	EyeColor  string   `json:"eye_color"`
	BirthYear string   `json:"birth_year"`
	Gender    string   `json:"gender"`
	Homeworld string   `json:"homeworld"`
	Films     []string `json:"films"`
	Species   []string `json:"species"`
	Vehicles  []string `json:"vehicles"`
	Starships []string `json:"starships"`
	Created   string   `json:"created"`
	Edited    string   `json:"edited"`
	URL       string   `json:"url"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	url := "https://swapi.dev/api/people/1"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var person Person
	
	err = json.Unmarshal(body, &person)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Fprintf(w, "Name: %s\n", person.Name)
	fmt.Fprintf(w, "Height: %s\n", person.Height)
	fmt.Fprintf(w, "Mass: %s\n", person.Mass)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	http.ListenAndServe(":8080", myRouter)
}

func main() {
	

	fmt.Println("Server listening on port 8080...")
	handleRequests()
	

	// fmt.Printf("Name: %s\n", person.Name)
	// fmt.Printf("Height: %s\n", person.Height)
	// fmt.Printf("Mass: %s\n", person.Mass)
}
