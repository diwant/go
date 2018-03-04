package main

import (
	"bufio"
	"bytes"
	"log"
	"math/rand"
	"os"
	"strings"
)

// Direction is N, E, S, W
type Direction uint8

const (
	_ = iota
	north
	east
	south
	west
)

// City sits adjacent to 0-4 cities in cardinal directions
type City struct {
	name          string              // No 2 Cities Share a Name
	neighbors     map[Direction]*City // [Cardinal Direction] => *City
	neighborIndex []Direction
	aliens        []*Alien
}

// NewCitiesFromFile Loads Cities From a Given File Name
func NewCitiesFromFile(fileName string) []*City {

	// Parse File and Read Into Cities Graph
	citiesFile, err := os.Open(fileName)

	// Exit on Error, Returning Empty Slice
	if err != nil {
		log.Println("Error while trying to load cities from", fileName, err)
		return []*City{}
	}

	// List of cities
	cities := []*City{}

	// Map of Cities For Quick Lookup When Adding Neighbors
	cityMap := make(map[string]*City) // [CityName] => *City

	// Map of City to Its' Neighbors
	cityNeighbors := make(map[string][]string) // [CityName] => ["dir=CityName", ...]

	// Iterate Over File
	citiesScanner := bufio.NewScanner(citiesFile)
	for citiesScanner.Scan() {

		// Read Next Scanned Line
		line := citiesScanner.Text()

		// Split Line By Spaces
		fields := strings.Split(line, " ")

		// Skip Empty Line
		if len(fields) < 1 {
			continue
		}

		// Create the City
		city := &City{
			name:      fields[0],
			neighbors: make(map[Direction]*City),
		}

		// Add to City List
		cities = append(cities, city)

		// Add to City Map
		cityMap[fields[0]] = city

		// Append Rest of Fields as Neighbors to City's Neighbor List
		cityNeighbors[fields[0]] = append([]string{}, fields[1:]...)
	}

	// Iterate Over Neighbors to Register Them to Cities
	for cityName, neighbors := range cityNeighbors {

		// Register Each Of Neighboring Cities
		for _, n := range neighbors {

			// Split NeighborInfo From Field By '='
			neighborInfo := strings.Split(n, "=")

			// Whoops, Incorrect Number of Info Fields
			if len(neighborInfo) != 2 {
				log.Printf("Received an incorrect number of fields for neighbor info for city, %s.  Expected 2 fields, received %d (from %s)\n", cityName, len(neighborInfo), neighborInfo)
			}

			// Parse Direction
			var neighborDir Direction

			switch neighborInfo[0] {
			case "north":
				neighborDir = north
			case "east":
				neighborDir = east
			case "south":
				neighborDir = south
			case "west":
				neighborDir = west
			}

			// Get Neighboring City
			neighborCity, found := cityMap[neighborInfo[1]]

			// If Neighbor Not Found, Create
			if !found {

				// Create the City
				neighborCity = &City{
					name:      neighborInfo[1],
					neighbors: make(map[Direction]*City),
				}

				// Add to City List
				cities = append(cities, neighborCity)

				// Add to City Map
				cityMap[neighborInfo[0]] = neighborCity
			}

			// Register Neighbor
			cityMap[cityName].RegisterNeighbor(neighborDir, neighborCity)
		}
	}

	return cities
}

// RegisterNeighbor registers a neighboring city in a given direction,
// returning false if a neighbor was already present in that direction
func (c *City) RegisterNeighbor(d Direction, neighbor *City) bool {

	// Search For Neighbor in Direction
	_, found := c.neighbors[d]
	if found {
		return false
	}

	// Add Neighbor to Map
	c.neighbors[d] = neighbor

	// Add direction to available index
	c.neighborIndex = append(c.neighborIndex, d)

	// Added New Neighbor, Return True
	return true
}

// DeregisterNeighbor deregisters a neighboring city,
// returning false if that neighbor was never present
func (c *City) DeregisterNeighbor(neighbor *City) bool {

	// Iterate Over Neighbors Map
	for dir, city := range c.neighbors {

		// City Found Among Neighbors?
		if city.name == neighbor.name {

			// Remove City From Neighbor Map
			delete(c.neighbors, dir)

			// Remove Direction From Neighbor Index
			for i, d := range c.neighborIndex {
				if d == dir {
					c.neighborIndex = append(c.neighborIndex[:i], c.neighborIndex[i+1:]...)
				}
			}

			// Work Done, Return
			return true
		}
	}

	// City Not Found, Return
	return false
}

// RegisterAlien adds alien to the list of aliens in the city
func (c *City) RegisterAlien(a *Alien) {

	// Append Alien
	c.aliens = append(c.aliens, a)
}

// DeregisterAlien adds alien to the list of aliens in the city
func (c *City) DeregisterAlien(a *Alien) {

	// Traverse Alien List and Remove Alien When Found
	for i, listAlien := range c.aliens {

		// Alien Found
		if listAlien.uuid == a.uuid {

			// Splice From List
			c.aliens = append(c.aliens[:i], c.aliens[i+1:]...)
			break
		}
	}
}

// GetRandomNeighbor Gets the Neighboring City in A Random Direction
func (c *City) GetRandomNeighbor() *City {

	// Return Nil for City if Isolated
	if len(c.neighborIndex) == 0 {
		return nil
	}

	// Get Index of Next City
	nextCityIndex := rand.Intn(len(c.neighborIndex))

	// Get Direction At that Index
	nextDirection := c.neighborIndex[nextCityIndex]

	// Get Next City
	return c.neighbors[nextDirection]
}

func (c *City) String() string {

	// Buffer to Compile City Info Into
	var buf bytes.Buffer

	// Print Name
	buf.WriteString(c.name)

	// Print Neighbors
	buf.WriteString("\n  Neighbors: ")
	for d, n := range c.neighbors {

		// Print Direction
		switch d {
		case north:
			buf.WriteString("north=")
		case east:
			buf.WriteString("east=")
		case south:
			buf.WriteString("south=")
		case west:
			buf.WriteString("west=")
		}

		// Print Name
		buf.WriteString(n.name)

		// Print Separator
		buf.WriteString(" ")
	}

	// Print Aliens
	buf.WriteString("\n Aliens: ")
	for _, a := range c.aliens {

		// Print Each Alien
		buf.WriteString(a.String())

		// Print Separator
		buf.WriteString(" ")
	}

	return buf.String()
}