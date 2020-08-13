package main

import (
	"math"
)

// x - cat A's position
// y - cat B's position
// z - mouse C's position
// who gets to the mouse first? If the same time, mouse escapes
// return "Cat A", "Cat B", or "Mouse C" on a new line
func catAndMouse(x int32, y int32, z int32) string {
	// x ----- z ----- y
	catADistance := math.Abs(float64(x - z))
	catBDistance := math.Abs(float64(y - z))

	if catADistance < catBDistance {
		return "Cat A"
	} else if catADistance > catBDistance {
		return "Cat B"
	} else {
		return "Mouse C"
	}
}
