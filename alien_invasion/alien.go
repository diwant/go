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

func (a *Alien) String() string {

	// Buffer to Render String To
	var buf bytes.Buffer

	// Print Alien Number and Num Moves and City (City to Make Sure It's Pointing To Same City It's Registered To)
	buf.WriteString(fmt.Sprintf("%d (%d) (%s)", a.uuid, a.numMoves, a.currCity.name))

	// Print 'x' if Dead
	if a.dead {
		buf.WriteString("x")
	}

	return buf.String()
}
