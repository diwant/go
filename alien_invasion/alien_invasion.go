package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {

	log.Println("Alien Invasion!!!")

	// Seed the Random Machine
	rand.Seed(time.Now().UTC().UnixNano())

	// Grab Command Line Args
	args := os.Args

	if len(args) < 3 {
		log.Fatalln("Usage: cities_file_name num_aliens")
		return
	}

	// Cities File Name
	citiesFileName := args[1]

	// Number of Aliens to Spread the Cities
	numAliens, err := strconv.ParseUint((args[2]), 10, 64)

	// Exit if Can't Parse Num Aliens
	if err != nil {
		log.Fatalln("Could not parse number of aliens from", args[2])
		return
	}

	// Confirm We Have Read the Arguments Correctly!
	log.Printf("City File Name: %s | Num Aliens: %d\n", citiesFileName, numAliens)

	// Create the Cities From the File
	cities := NewCitiesFromFile(citiesFileName)

	// Exit if 0 Cities
	if len(cities) == 0 {
		log.Fatalln("No cities loaded")
		return
	}

	// Hold the Aliens in This Slice
	aliens := []*Alien{}

	// Deploy the N Aliens
	var n uint64
	for n = 0; n < numAliens; n++ {

		// Pick a Random City Index
		cityIndex := rand.Intn(len(cities))

		// Deploy this Alien in That City
		a := NewAlien(n, cities[cityIndex])

		// Append Alien To Our Slice of Aliens
		aliens = append(aliens, a)
	}

	// Debug Print
	for _, city := range cities {
		fmt.Println(city)
	}

	// Start the Game Loop
	for {

		// Iterate Over Cities and Run Alien Battles
		for _, c := range cities {

			// If City Destroyed, Move On
			if c.destroyed {
				continue
			}

			// More Than One Alien?
			if len(c.aliens) > 1 {

				// Aliens Fought and City is Destroyed
				log.Println(c.Explode())
			}
		}

		// Flag If Every Alive Alien Has Moved > 10000 Times
		movesOver := true

		// Prune Dead Aliens, Migrate Others
		for i := len(aliens) - 1; i >= 0; i-- {

			// Is Alien Dead?
			if aliens[i].dead {

				// Prune From Aliens
				aliens = append(aliens[:i], aliens[i+1:]...)
				continue
			}

			// Alien Not Dead, Migrate
			if aliens[i].Travel() <= 10000 {

				// Found An Alien That Has Moves Left
				movesOver = false
			}
		}

		// All Aliens Have Been Destroyed?
		if len(aliens) == 0 {
			break
		}

		// No Alive Alien Has Moves Left?
		if movesOver {
			break
		}
	}

	// Debug Print
	for _, city := range cities {
		fmt.Println(city)
	}

	log.Println("...and so ran the alien invasion of 2018.  THE END")
}
