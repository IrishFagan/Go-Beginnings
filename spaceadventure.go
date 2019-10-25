package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"time"
	"os"
)

//Creates struct for Planets
type Planet struct {
	Name string `json: "name"`
	Description string `json: "description"`
}

//Creates list of Planets containing their name and description
type Planets struct {
	Planets []Planet `json: "planets"`
}

//Calls conversation()
func main() {
	conversation()
}

//Used to retrieve name of the user, then calls handleInput while passing it
//the names and descriptions of the planets
func conversation() {	
	file, _ := ioutil.ReadFile("planetarySystem.json")
	p_data := Planets{}
	_ = json.Unmarshal([]byte(file), &p_data)

	var nameInput string
	fmt.Println("Welcome to our quaint solar system!")
	fmt.Println("There used to be 9 planets but now there are only 8...")
	fmt.Println("...")
	fmt.Println("..anyways, what would you prefer I call you?\n")
	fmt.Println("ENTER NAME PLS: ")
	fmt.Scanln(&nameInput)
	fmt.Println("\nHmm It's a pleasure to meet you, " + nameInput)
	fmt.Println("My name is Gort, and I'm here to show you around our solar system!")

	handleInput(p_data)
}

//Handles the input from the user of whether they would like to hear about a
//random or specific planet
func handleInput(p_data Planets) {
	var randPlanet string
	fmt.Println("Would you like to hear about a random planet?")
	fmt.Println("Y or N")
	fmt.Scanln(&randPlanet)

	if randPlanet == "Y" {
		rand.Seed(time.Now().UnixNano())
		var randInt = rand.Intn(len(p_data.Planets))
		fmt.Println("Name: " + p_data.Planets[randInt].Name)
		fmt.Println("Desc: " + p_data.Planets[randInt].Description)
		os.Exit(3)
	} else if randPlanet == "N" {
		fmt.Println("Alright, what planet would you like to hear about then?")
		specificPlanet(p_data)
	} else if randPlanet != "Y" {
		handleInput(p_data)
	} else if randPlanet != "N" {
		handleInput(p_data)
	}
}

//Handles the input of a specific planet from the user
func specificPlanet(p_data Planets) {
	var planetInput string
	var planetFound bool
	
	fmt.Scanln(&planetInput)

	for i := 0; i < len(p_data.Planets); i++ {
		if planetInput == p_data.Planets[i].Name {
			fmt.Println(planetInput)
			fmt.Println(p_data.Planets[i].Description)
			planetFound = true
		}
	}

	if planetFound != true {
		fmt.Println("Sorry, I don't know that planet!")
		fmt.Println("Again, what planet would you like to hear about?")
		specificPlanet(p_data)
	}
}