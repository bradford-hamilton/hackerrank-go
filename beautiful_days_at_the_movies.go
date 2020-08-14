package main

import "strconv"

func beautifulDays(firstDay int32, lastDay int32, k int32) int32 {
	var bDays int32
	for i := firstDay; i <= lastDay; i++ {
		orig := i
		rev := reverse(i)
		if (orig-rev)%k == 0 {
			bDays++
		}
	}
	return bDays
}

func reverse(ind int32) int32 {
	str := strconv.Itoa(int(ind))
	var reversedStr string
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}
	rev, _ := strconv.Atoi(reversedStr)
	return int32(rev)
}
