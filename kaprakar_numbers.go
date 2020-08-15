package main

import (
	"fmt"
	"strconv"
)

// https://www.hackerrank.com/challenges/kaprekar-numbers/problem
// Complete the kaprekarNumbers function below.
// p - lower limit
// q - higher limit
func kaprekarNumbers(p int32, q int32) {
	for i := p; i <= q; i++ {
		sq := i * i
		str := strconv.Itoa(int(sq))
		if len(str) == 1 {
			continue
		}
		num1, num2 := splitNums(str)
		if num1+num2 == i {
			fmt.Println(i)
		}
	}
}

func splitNums(num string) (int32, int32) {
	num1 := num[:len(num)/2]
	num2 := num[len(num)/2:]

	n1, _ := strconv.Atoi(num1)
	n2, _ := strconv.Atoi(num2)

	return int32(n1), int32(n2)
}
