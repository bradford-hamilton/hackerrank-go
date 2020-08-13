package main

import "fmt"

// https://www.hackerrank.com/challenges/bon-appetit/problem
// bill: an array of integers representing the cost of each item ordered
// k: an integer representing the zero-based index of the item Anna doesn't eat
// b: the amount of money that Anna contributed to the bill
// If Brian did not overcharge Anna, print Bon Appetit on a new line; otherwise,
// print the difference (i.e., bCharged - bActual) that Brian must refund to Anna.
// This will always be an integer.
func bonAppetit(bill []int32, k int32, b int32) {
	var sum int32
	for _, v := range bill {
		sum += v
	}
	annaSum := sum - bill[k]
	annaShouldHaveContributed := annaSum / 2
	annaContributed := b

	if annaShouldHaveContributed >= annaContributed {
		fmt.Println("Bon Appetit")
	} else {
		fmt.Println(annaContributed - annaShouldHaveContributed)
	}
}
