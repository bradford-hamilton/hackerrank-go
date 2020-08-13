package main

// https://www.hackerrank.com/challenges/countingsort2/problem
// counts is an array of how may times each number from arr index was used.
// original array: [1, 2, 1, 2, 5, 6]
// counts array:   [2, 2, 2, 2, 1, 1]
func countingSort2(arr []int32) []int32 {
	counts := make([]int32, 100)
	for _, v := range arr {
		counts[v]++
	}

	res := make([]int32, 100)
	for i := range counts {
		counts[i+1] = counts[i] + counts[i+1]
	}

	for i := range arr {
		ind := arr[i]
		counts[ind]--
		indexCount := counts[ind]
		res[indexCount] = ind
	}

	return res
}
