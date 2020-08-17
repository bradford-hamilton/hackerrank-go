package main

import "math"

// https://www.hackerrank.com/challenges/sherlock-and-squares/problem
func squares(a int32, b int32) int32 {
	var numSquares int32
	for i := a; i <= b; i++ {
		sr := math.Sqrt(float64(i))
		if sr == math.Floor(sr) {
			numSquares++
		}
	}
	return numSquares
}
