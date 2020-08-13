package main

// https://www.hackerrank.com/challenges/insertionsort2/problem
func insertionSort2(n int32, arr []int32) {
	for i := int32(1); i < n; i++ {
		v := arr[i]
		j := i

		for j > 0 && arr[j-1] > v {
			arr[j] = arr[j-1]
			j--
		}

		arr[j] = v
		printArr(arr)
	}
}
