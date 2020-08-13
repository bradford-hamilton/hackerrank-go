package main

import "math"

// https://www.hackerrank.com/challenges/drawing-book/problem?h_r=next-challenge&h_v=zen&h_r=next-challenge&h_v=zen
// n - number of pages in the book
// p - the page we need to flip to
// Calculate the least amount of page turns to get to that page.
// You could be starting from the front or the back.
func pageCount(n int32, p int32) int32 {
	pageTurns := float64(p) / 2
	totalTurns := float64(n) / 2
	return int32(math.Min(pageTurns, totalTurns-pageTurns))
}

// Note: 1 of 27 tests still fails with the above code...
