package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateIsolatedCity(t *testing.T) {

	// Create Isolated City From Scenario
	cities := NewCitiesFromFile("scenarios/isolated_city.txt")

	// Expect Only 1 City
	assert.Equal(t, len(cities), 1, "Unexpected number of cities")

	// Get Created City
	c := cities[0]

	// Expect Name is Foo
	assert.Equal(t, "Foo", c.name, "Unexpected city name")

	// Expect 0 Neighbors
	assert.Equal(t, 0, len(c.neighborIndex), "Unexpected number of neighbors in the neighbor index")
	assert.Equal(t, 0, len(c.neighbors), "Unexpected number of neighbors")

	// Expect Zero Aliens
	assert.Equal(t, 0, len(c.aliens), "Unexpected number of resident aliens")
}

func TestCreateTwoCitiesBothLink(t *testing.T) {

	// Create Isolated City From Scenario
	cities := NewCitiesFromFile("scenarios/2_cities-both_link.txt")

	// Expect Only 2 Cities
	assert.Equal(t, len(cities), 2, "Unexpected number of cities")

	// Check Each Created City
	for i, c := range cities {

		switch i {
		case 0:

			// Expect Name is Foo
			assert.Equal(t, "Foo", c.name, "Unexpected city name")

			// Expect 1 Neighbors
			assert.Equal(t, 1, len(c.neighborIndex), "Unexpected number of neighbors in the neighbor index")
			assert.Equal(t, 1, len(c.neighbors), "Unexpected number of neighbors")

			// Expect Zero Aliens
			assert.Equal(t, 0, len(c.aliens), "Unexpected number of resident aliens")
		case 1:

			// Expect Name is Bar
			assert.Equal(t, "Bar", c.name, "Unexpected city name")

			// Expect 1 Neighbors
			assert.Equal(t, 1, len(c.neighborIndex), "Unexpected number of neighbors in the neighbor index")
			assert.Equal(t, 1, len(c.neighbors), "Unexpected number of neighbors")

			// Expect Zero Aliens
			assert.Equal(t, 0, len(c.aliens), "Unexpected number of resident aliens")
		}
	}
}

func TestCreateTwoCitiesOneLinks(t *testing.T) {

	// Create Isolated City From Scenario
	cities := NewCitiesFromFile("scenarios/2_cities-one_links.txt")

	// Expect Only 2 Cities
	assert.Equal(t, 2, len(cities), "Unexpected number of cities")

	// Check Each Created City
	for i, c := range cities {

		switch i {
		case 0:

			// Expect Name is Foo
			assert.Equal(t, "Foo", c.name, "Unexpected city name")

			// Expect 1 Neighbors
			assert.Equal(t, 1, len(c.neighborIndex), "Unexpected number of neighbors in the neighbor index")
			assert.Equal(t, 1, len(c.neighbors), "Unexpected number of neighbors")

			// Expect Zero Aliens
			assert.Equal(t, 0, len(c.aliens), "Unexpected number of resident aliens")
		case 1:

			// Expect Name is Bar
			assert.Equal(t, "Bar", c.name, "Unexpected city name")

			// Expect 0 Neighbors
			assert.Equal(t, 0, len(c.neighborIndex), "Unexpected number of neighbors in the neighbor index")
			assert.Equal(t, 0, len(c.neighbors), "Unexpected number of neighbors")

			// Expect Zero Aliens
			assert.Equal(t, 0, len(c.aliens), "Unexpected number of resident aliens")
		}
	}
}

func TestBattleInLinkedCity(t *testing.T) {

	// Create Isolated City From Scenario
	cities := NewCitiesFromFile("scenarios/2_cities-both_link.txt")

	// Create a Left and Right Alien For the Battle In The Same City
	leftAlien := NewAlien(0, cities[0])
	rightAlien := NewAlien(1, cities[0])

	// Explode City
	explodeMsg := cities[0].Explode()

	// Set Up Expected Explode Message
	expectedExplodeMsg := fmt.Sprintf("%s has been destroyed by alien %d and alien %d\n", cities[0].name, leftAlien.uuid, rightAlien.uuid)

	// Assert the Explode Message is Correct
	assert.Equal(t, expectedExplodeMsg, explodeMsg, "Unexpected explode message")

	// Assert Neighboring City Doesn't Link In
	assert.Equal(t, 0, len(cities[1].neighborIndex), "Unexpected number of neighbors in index list")
	assert.Equal(t, 0, len(cities[1].neighbors), "Unexpected number of neighbors in map")

	// Assert Aliens Are Dead
	assert.Equal(t, true, leftAlien.dead, "Left Alien Should Be Dead")
	assert.Equal(t, true, rightAlien.dead, "Right Alien Should Be Dead")
}

func TestNonBattleInLinkedCity(t *testing.T) {

	// Create Isolated City From Scenario
	cities := NewCitiesFromFile("scenarios/2_cities-both_link.txt")

	// Create a Left Alien In the First City
	leftAlien := NewAlien(0, cities[0])

	// Create a Right Alien In the Second City
	rightAlien := NewAlien(1, cities[1])

	// Explode City
	explodeMsg := cities[0].Explode()

	// Set Up Expected Explode Message
	expectedExplodeMsg := fmt.Sprintf("%s has been destroyed", cities[0].name)

	// Assert the Explode Message is Correct
	assert.Equal(t, expectedExplodeMsg, explodeMsg, "Unexpected explode message")

	// Assert Neighboring City Doesn't Link In
	assert.Equal(t, 0, len(cities[1].neighborIndex), "Unexpected number of neighbors in index list")
	assert.Equal(t, 0, len(cities[1].neighbors), "Unexpected number of neighbors in map")

	// Assert Left Alien Is Dead
	assert.Equal(t, true, leftAlien.dead, "Left Alien Should Be Dead")

	// Assert Right Alien Is NOT Dead (Because It Was In A Different City!)
	assert.Equal(t, false, rightAlien.dead, "Right Alien Should Be Alive")
}
