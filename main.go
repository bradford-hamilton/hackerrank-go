package main

import "fmt"

func printArr(arr []int32) {
	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
	fmt.Print("\n")
}

func main() {
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	found := binarySearch(items, 55)
	fmt.Println(found)
}
