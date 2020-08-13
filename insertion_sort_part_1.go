package main

// https://www.hackerrank.com/challenges/insertionsort1/problem
func insertionSort1(n int32, arr []int32) {
	i := n - 1
	j := i
	v := arr[i]

	for j > 0 && arr[j-1] > v {
		arr[j] = arr[j-1]
		j--
		printArr(arr)
	}
	arr[j] = v
	printArr(arr)
}
