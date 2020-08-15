package main

import "math"

// https://www.hackerrank.com/challenges/minimum-distances/problem
// [7, 1, 3, 4, 1, 7]
func minimumDistances(a []int32) int32 {
	m := make(map[int32][]int32)
	for i := 0; i < len(a); i++ {
		m[a[i]] = append(m[a[i]], int32(i))
	}

	minDistances := getMinDistances(m)
	if len(minDistances) == 0 {
		return -1
	}

	return smallest(int32(len(a)), minDistances)
}

func getMinDistances(m map[int32][]int32) []int32 {
	var minDistances []int32
	for _, v := range m {
		if len(v) > 1 {
			minDistances = append(minDistances, getMinDistance(v))
		}
	}
	return minDistances
}

// get the minimum distance between the provided indexes
func getMinDistance(distances []int32) int32 {
	if len(distances) <= 2 {
		return int32(math.Abs(float64(distances[0] - distances[1])))
	}

	smallest := math.Abs(float64(distances[0] - distances[1]))
	for i := 0; i < len(distances); i++ {
		if i == len(distances)-1 {
			break
		}
		if math.Abs(float64(distances[i]-distances[i+1])) < smallest {
			smallest = math.Abs(float64(distances[i] - distances[i+1]))
		}
	}

	return int32(smallest)
}

func smallest(arrayLen int32, minDistances []int32) int32 {
	smallest := arrayLen
	for _, v := range minDistances {
		if v < smallest {
			smallest = v
		}
	}
	return smallest
}
