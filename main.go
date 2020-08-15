package main

import "fmt"

func printArr(arr []int32) {
	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
	fmt.Print("\n")
}

func main() {
	res := repeatedString("aba", 1000000000000)
	fmt.Println(res)
}
