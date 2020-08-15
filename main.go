package main

import "fmt"

func printArr(arr []int32) {
	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
	fmt.Print("\n")
}

func main() {
	minimumDistances([]int32{7, 1, 3, 4, 1, 7})
}
