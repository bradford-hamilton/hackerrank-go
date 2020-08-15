package main

import "strconv"

// https://www.hackerrank.com/challenges/find-digits/problem
func findDigits(n int32) int32 {
	var count int32
	for _, v := range strconv.Itoa(int(n)) {
		num, _ := strconv.Atoi(string(v))
		if num == 0 {
			continue
		}
		if n%int32(num) == 0 {
			count++
		}
	}
	return count
}
