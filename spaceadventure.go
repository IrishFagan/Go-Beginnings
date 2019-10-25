package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"time"
	"os"
)

type Planet struct {
	Name string `json: "name"`
	Description string `json: "description"`
}

type Planets struct {
	Planets []Planet `json: "planets"`
}

func main() {
	conversation()
}

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
		specificPlanet(p_data)
	} else if randPlanet != "Y" {
		handleInput(p_data)
	} else if randPlanet != "N" {
		handleInput(p_data)
	}
}

func specificPlanet(p_data Planets) {
	var planetInput string
	var planetFound bool
	
	fmt.Println("Alright, then what planet would you like to hear about?")
	fmt.Scanln(&planetInput)

	for i := 0; i < len(p_data.Planets); i++ {
		if planetInput == p_data.Planets[i].Name {
			fmt.Println(planetInput)
			fmt.Println(p_data.Planets[i].Description)
			planetFound = true
		}
	}

	if planetFound != true {
		fmt.Print("Sorry, I don't know that planet!")
		specificPlanet(p_data)
	}
}