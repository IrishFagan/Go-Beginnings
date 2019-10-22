package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"time"
)

type Planets struct {
	Planets []Planet `json: "planets"`
}


type Planet struct {
	Name string `json: "name"`
	Description string `json: "description"`
}

func main() {
	file, _ := ioutil.ReadFile("planetarySystem.json")
	data := Planets{}

	_ = json.Unmarshal([]byte(file), &data)

	fmt.Println("Welcome to our quaint solar system!")
	fmt.Println("There used to be 9 planets but now there are only 8...")
	fmt.Println("...")
	fmt.Println("..anyways, what would you prefer I call you?\n")
	fmt.Println("ENTER NAME PLS: ")
	var nameInput string
	fmt.Scanln(&nameInput)
	fmt.Println("\nHmm It's a pleasure to meet you, " + nameInput)
	fmt.Println("My name is Gort, and I'm here to show you around our solar system!")
	var randPlanet string
	fmt.Println("Would you like to hear about a random planet?")
	fmt.Println("Y or N")
	fmt.Scanln(&randPlanet)

	for {
		if randPlanet == "Y" {
			rand.Seed(time.Now().UnixNano())
			var randInt = rand.Intn(len(data.Planets))
			fmt.Println(randInt)
			fmt.Println("Name: " + data.Planets[randInt].Name)
			fmt.Println("Desc: " + data.Planets[randInt].Description)
			break
		} else if randPlanet == "N" {
			break
		} else if randPlanet != "Y" {
			fmt.Println("I didn't catch that.")
			fmt.Println("Would you like to hear about a random planet?")
			fmt.Scanln(&randPlanet)
		} else if randPlanet != "N" {
			fmt.Println("I didn't catch that.")
			fmt.Println("Would you like to hear about a random planet?")
			fmt.Scanln(&randPlanet)
		}
	}

	if randPlanet == "N" {
		fmt.Println("Alright, then what planet would you like to hear about?")
	}

	var planetInput string
	fmt.Scanln(&planetInput)
	for i := 0; i < len(data.Planets); i++ {
		if planetInput == data.Planets[i].Name {
			fmt.Println(planetInput)
			fmt.Println(data.Planets[i].Description)
		} else {
			fmt.Println("Sorry, I don't know that planet...")
			break
		}
	}

	fmt.Print("\n")
}