package main

// https://www.hackerrank.com/challenges/migratory-birds/problem
// Each time a bird is spotted, it's ID is placed in the arr
// Print most common bird
// If equally common, choose smaller ID
func migratoryBirds(arr []int32) int32 {
	m := make(map[int32]int32)
	for _, v := range arr {
		m[v]++
	}

	var mostSightedID int32
	var mostSights int32

	for birdID, numOfSights := range m {
		if numOfSights > mostSights {
			mostSights = numOfSights
			mostSightedID = birdID
		} else if numOfSights == mostSights {
			if birdID < mostSightedID {
				mostSights = numOfSights
				mostSightedID = birdID
			}
		}
	}

	return mostSightedID
}
