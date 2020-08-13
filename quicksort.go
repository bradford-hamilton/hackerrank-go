package main

import "fmt"

// https://www.hackerrank.com/challenges/quicksort1/problem
func quickSort(arr []int32) []int32 {
	var left, equal, right []int32
	pivot := arr[0]

	for _, v := range arr {
		if v < pivot {
			left = append(left, v)
		}
		if v == pivot {
			equal = append(equal, v)
		}
		if v > pivot {
			right = append(right, v)
		}
	}

	var res []int32
	res = append(res, left...)
	res = append(res, equal...)
	res = append(res, right...)

	for _, v := range res {
		fmt.Println(v)
	}

	return res
}
