package main

import "fmt"

// s: integer, starting point of Sam's house location.
// t: integer, ending location of Sam's house location.
// a: integer, location of the Apple tree.
// b: integer, location of the Orange tree.
// apples: integer array, distances at which each apple falls from the tree.
// oranges: integer array, distances at which each orange falls from the tree.

// Complete the countApplesAndOranges function below.
func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	var houseApples int
	for _, v := range apples {
		if (a+v >= s) && (a+v <= t) {
			houseApples++
		}
	}

	var houseOranges int
	for _, v := range oranges {
		if (b+v <= t) && (b+v >= s) {
			houseOranges++
		}
	}

	fmt.Println(houseApples)
	fmt.Println(houseOranges)
}
