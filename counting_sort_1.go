package main

// https://www.hackerrank.com/challenges/countingsort2/problem
func countingSort(arr []int32) []int32 {
	// all nums fit within 0-100
	res := make([]int32, 100)
	for _, v := range arr {
		res[v]++
	}
	return res
}

// [63 25 73 1 98 73 56 84 86 57 16 83]
// [0 2 0 2 0 0 1 0 1 2 1 0 1 1 0 0 2 ]
