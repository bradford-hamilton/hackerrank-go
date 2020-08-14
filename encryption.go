package main

import (
	"math"
)

// https://www.hackerrank.com/domains/algorithms?filters%5Bstatus%5D%5B%5D=solved
func encryption(s string) string {
	sLen := len(s)
	srFloat := math.Sqrt(float64(sLen))
	jump := int(srFloat)
	if jump*jump != sLen {
		jump++
	}

	var ret string
	for i := 0; i < jump; i++ {
		for j := i; j < sLen; j += jump {
			ret += string(s[j])
		}
		ret += " "
	}

	return ret
}
