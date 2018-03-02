package main

import (
	"fmt"
	"os"
	"strconv"
)

// City sits adjacent to 0-4 cities in cardinal directions
type City struct {
	name                     string
	north, east, south, west *City
}

func main() {

	fmt.Println("Alien Invasion!!!")

	// Grab Command Line Args
	args := os.Args

	if len(args) < 3 {
		fmt.Println("Usage: cities_file_name num_aliens")
		return
	}

	// Cities File Name
	citiesFileName := args[1]

	// Number of Aliens to Spread the Cities
	numAliens, err := strconv.Atoi(args[2])

	// Exit if Can't Parse Num Aliens
	if err != nil {
		fmt.Println("Could not parse number of aliens from", args[2])
		return
	}

	// Confirm We Have Read the Arguments Correctly!
	fmt.Printf("City File Name: %s | Num Aliens: %d\n", citiesFileName, numAliens)

	// Parse File and Read Into Cities Graph
}
