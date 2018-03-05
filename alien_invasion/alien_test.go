package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrappedAlienTravels(t *testing.T) {

	// Create Isolated City
	cities := NewCitiesFromFile("scenarios/isolated_city.txt")

	// Create a New Alien
	a := NewAlien(1, cities[0])

	// Assert Starts with 0 Moves
	assert.Equal(t, uint16(0), a.numMoves, "Unexpected number of moves")

	// Travel
	a.Travel()

	// Assert Moves Increments
	assert.Equal(t, uint16(1), a.numMoves, "Unexpected number of moves")

	// Assert Same City
	assert.Equal(t, "Foo", a.currCity.name, "Unexpected city name for current location")

	// Assert Population Count
	assert.Equal(t, 1, len(cities[0].aliens), "Unexpected number of resident aliens")
}

func TestNontrappedAlienTravelsInBothLink(t *testing.T) {

	// Create Isolated City
	cities := NewCitiesFromFile("scenarios/2_cities-both_link.txt")

	// Create a New Alien
	a := NewAlien(1, cities[0])

	// Assert Starts with 0 Moves
	assert.Equal(t, uint16(0), a.numMoves, "Unexpected number of moves")

	// Assert Population Count
	assert.Equal(t, 1, len(cities[0].aliens), "Unexpected number of resident aliens")
	assert.Equal(t, 0, len(cities[1].aliens), "Unexpected number of resident aliens")

	// Travel
	a.Travel()

	// Assert Moves Increments
	assert.Equal(t, uint16(1), a.numMoves, "Unexpected number of moves")

	// Assert Second City
	assert.Equal(t, "Bar", a.currCity.name, "Unexpected city name for current location")

	// Assert Population Count
	assert.Equal(t, 0, len(cities[0].aliens), "Unexpected number of resident aliens")
	assert.Equal(t, 1, len(cities[1].aliens), "Unexpected number of resident aliens")

	// Travel
	a.Travel()

	// Assert Moves Increments
	assert.Equal(t, uint16(2), a.numMoves, "Unexpected number of moves")

	// Assert In First City
	assert.Equal(t, "Foo", a.currCity.name, "Unexpected city name for current location")

	// Assert Population Count
	assert.Equal(t, 1, len(cities[0].aliens), "Unexpected number of resident aliens")
	assert.Equal(t, 0, len(cities[1].aliens), "Unexpected number of resident aliens")

	// Travel
	a.Travel()

	// Assert Moves Increments
	assert.Equal(t, uint16(3), a.numMoves, "Unexpected number of moves")

	// Assert In Second City
	assert.Equal(t, "Bar", a.currCity.name, "Unexpected city name for current location")

	// Assert Population Count
	assert.Equal(t, 0, len(cities[0].aliens), "Unexpected number of resident aliens")
	assert.Equal(t, 1, len(cities[1].aliens), "Unexpected number of resident aliens")
}

func TestNontrappedAlienTravelsInOneLinks(t *testing.T) {

	// Create Isolated City
	cities := NewCitiesFromFile("scenarios/2_cities-one_links.txt")

	// Create a New Alien
	a := NewAlien(1, cities[0])

	// Assert Starts with 0 Moves
	assert.Equal(t, uint16(0), a.numMoves, "Unexpected number of moves")

	// Assert Population Count
	assert.Equal(t, 1, len(cities[0].aliens), "Unexpected number of resident aliens")
	assert.Equal(t, 0, len(cities[1].aliens), "Unexpected number of resident aliens")

	// Travel
	a.Travel()

	// Assert Moves Increments
	assert.Equal(t, uint16(1), a.numMoves, "Unexpected number of moves")

	// Assert Second City
	assert.Equal(t, "Bar", a.currCity.name, "Unexpected city name for current location")

	// Assert Population Count
	assert.Equal(t, 0, len(cities[0].aliens), "Unexpected number of resident aliens")
	assert.Equal(t, 1, len(cities[1].aliens), "Unexpected number of resident aliens")

	// Travel - Now Trapped in a City That Doesn't Link Out
	a.Travel()

	// Assert Moves Increments
	assert.Equal(t, uint16(2), a.numMoves, "Unexpected number of moves")

	// Assert Second City Still
	assert.Equal(t, "Bar", a.currCity.name, "Unexpected city name for current location")

	// Assert Population Count
	assert.Equal(t, 0, len(cities[0].aliens), "Unexpected number of resident aliens")
	assert.Equal(t, 1, len(cities[1].aliens), "Unexpected number of resident aliens")
}
