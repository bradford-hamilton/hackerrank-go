package main

import "math"

// https://www.hackerrank.com/challenges/sock-merchant/problem?h_r=next-challenge&h_v=zen
// It must return an integer representing the number of matching pairs of socks that are available.
// n: the number of socks in the pile
// ar: the colors of each sock
func sockMerchant(n int32, ar []int32) int32 {
	m := make(map[int32]int32)
	for _, v := range ar {
		m[v]++
	}

	var numPairs int32
	for _, numSocks := range m {
		numPairs += int32(math.Floor(float64(numSocks) / 2))
	}

	return numPairs
}
