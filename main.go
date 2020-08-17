package main

import "fmt"

func printArr(arr []int32) {
	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
	fmt.Print("\n")
}

func main() {
	ok := isPalindrome("abababa")
	fmt.Println(ok)
}
