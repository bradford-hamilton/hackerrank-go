package main

// https://www.hackerrank.com/challenges/divisible-sum-pairs/problem
// You are given array of ints, and a positive integer - k
// Find and print the number of (i, j) pairs where i < j and arr[i] + arr[j] is divisible by 5
func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	var numPairs int32
	for i := int32(0); i < n-1; i++ {
		for j := int32(0); j < n; j++ {
			if (ar[i]+ar[j])%k == 0 && i < j {
				numPairs++
			}
		}
	}
	return numPairs
}
