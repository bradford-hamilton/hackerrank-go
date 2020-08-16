package main

import "math"

// https://www.hackerrank.com/challenges/halloween-sale/problem?h_r=next-challenge&h_v=zen
// Return the number of games you can buy
// first game you buy is at `p` dollars
// every subsequent must be `d` dollars less than the cost of the previous one bought.
// This will continue until the cost becomes less than or equal to m dollars, after which every game you buy will cost m dollars each.
func howManyGames(p int32, d int32, m int32, s int32) int32 {
	var count int32
	for s >= p {
		count++
		s -= p
		p = int32(math.Max(float64(p-d), float64(m)))
	}
	return count
}
