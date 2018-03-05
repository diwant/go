package main

import (
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
