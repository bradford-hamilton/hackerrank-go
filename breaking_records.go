package main

// https://www.hackerrank.com/challenges/breaking-best-and-worst-records/problem
func breakingRecords(scores []int32) []int32 {
	high := scores[0]
	low := scores[0]
	var highCount int32 = 1
	var lowCount int32 = 1

	for _, v := range scores {
		if v > high {
			high = v
			highCount++
		} else if v < low {
			low = v
			lowCount++
		}
	}

	return []int32{highCount - 1, lowCount - 1}
}
