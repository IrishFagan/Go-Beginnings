package main

import (
	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"
	"math/rand"
)

type Planets struct {
	Planets []Planet `json:"planets"`
}

type Planet struct {
	name 		string `json:"name"`
	description string `json:"description"`
}

func main() {
	jsonFile, err := os.Open("planetarySystem.json")

	if err != nil {
		fmt.Println(err)	
	}
	fmt.Println("Opened planetarySystem.json!")

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var planets Planets
	json.Unmarshal(byteValue, &planets)

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
			var randInt = rand.Intn(len(planets.Planets))
			fmt.Println(randInt)
			fmt.Println("Name: " + planets.Planets[randInt].name)
			fmt.Println("Desc: " + planets.Planets[randInt].description)
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

	//var planetInput string
	//if planetInput not one of nine options {
	//	Ask user to repeat themselves
	//	fmt.Scanln(&planetInput)
	//}

	//if planetInput is one of nine options {
	//	Print name and description of planet
	//}

	fmt.Print("\n")
}