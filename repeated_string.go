package main

import (
	"math"
	"strings"
)

// This seems to solve the problem in some of the cases, some are still not passing.
// I'd like to figure out how to count without building the string at all.
// https://www.hackerrank.com/challenges/repeated-string/problem
func repeatedString(s string, n int64) int64 {
	sLen := len(s)
	wantLen := int(n)
	numCharsDiff := wantLen - sLen
	numRepeats := math.Ceil(float64(numCharsDiff) / float64(len(s)))
	newStr := strings.Repeat(s, int(numRepeats))
	newStr = newStr[:wantLen-1]
	finalStr := newStr + string(s[len(s)-1])
	return countLetterA(finalStr)
}

func countLetterA(str string) int64 {
	var count int64
	for _, v := range str {
		if v == 'a' {
			count++
		}
	}
	return count
}
