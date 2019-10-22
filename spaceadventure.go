package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"time"
	"os"
)

//JSON Structs
type Planets struct {
	Planets []Planet `json: "planets"`
}


type Planet struct {
	Name string `json: "name"`
	Description string `json: "description"`
}

//Only function(I realize how ugly this looks but I felt too deep to reformat into indivindual functions. Will go back and fix
//so that it's not as ugly)
func main() {
	file, _ := ioutil.ReadFile("planetarySystem.json")
	data := Planets{}

	_ = json.Unmarshal([]byte(file), &data)

	//Ask for name and if user would like a random planet
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

	//take input Y or N
	for {
		if randPlanet == "Y" {
			rand.Seed(time.Now().UnixNano())
			var randInt = rand.Intn(len(data.Planets))
			fmt.Println(randInt)
			fmt.Println("Name: " + data.Planets[randInt].Name)
			fmt.Println("Desc: " + data.Planets[randInt].Description)
			os.Exit(3)
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

	//handle input N
	if randPlanet == "N" {
		fmt.Println("Alright, then what planet would you like to hear about?")
	}
	var planetInput string
	var planetFound bool
	fmt.Scanln(&planetInput)

	for i := 0; i < len(data.Planets); i++ {
		if planetInput == data.Planets[i].Name {
			fmt.Println(planetInput)
			fmt.Println(data.Planets[i].Description)
			planetFound = true
		}
	}

	if planetFound != true {
		fmt.Print("Sorry, I don't know that planet!")
	}

	//program doesn't ask again if input doesn't match acutal planet. Need to reformat when I actually
	//reformat into indivindual functions since that will make the implementation a lot less messy

	fmt.Print("\n")
}