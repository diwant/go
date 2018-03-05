package main

import (
	"bytes"
	"fmt"
)

// Alien Visits Cities And Fights Other Aliens
type Alien struct {
	uuid     uint64 // Could have a HUGE number of aliens
	currCity *City
	dead     bool
	numMoves uint16
}

// NewAlien Creates an Alien and Registers it to the Given City
func NewAlien(uuid uint64, c *City) *Alien {
	a := &Alien{
		uuid:     uuid,
		currCity: c,
	}
	c.RegisterAlien(a)
	return a
}

// Travel is When An Alien Leaves One City For Another
func (a *Alien) Travel() {

	// Increment Number of Moves
	a.numMoves++

	// Get Next City Travelling in an Random Direction
	nextCity := a.currCity.GetRandomNeighbor()

	// Is Alien Trapped?
	if nextCity == nil {
		return
	}

	// Deregister from Current City
	a.currCity.DeregisterAlien(a)

	// Switch City
	a.currCity = nextCity

	// Register In New City
	a.currCity.RegisterAlien(a)

}

// String Renders an Alien as a String
func (a *Alien) String() string {

	// Buffer to Render String To
	var buf bytes.Buffer

	// Print Alien Number and Num Moves
	buf.WriteString(fmt.Sprintf("%d (%d)", a.uuid, a.numMoves))

	// Print City, If Present
	if a.currCity != nil {
		buf.WriteString(fmt.Sprintf(" [%s]", a.currCity.name))
	}

	// Print 'x' if Dead
	if a.dead {
		buf.WriteString("x")
	}

	return buf.String()
}
