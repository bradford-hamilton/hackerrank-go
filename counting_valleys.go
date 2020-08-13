package main

// https://www.hackerrank.com/challenges/counting-valleys/problem?h_r=next-challenge&h_v=zen&h_r=next-challenge&h_v=zen&h_r=next-challenge&h_v=zen
// n - total number of steps in gary's hike
// s - UUDDUDDDUUUUUUUDDUDDD - U: upwords step, D: downwards step
func countingValleys(n int32, s string) int32 {
	var hike int32
	var mountains int32
	var valleys int32

	for _, step := range s {
		if hike == 0 && step == 'U' {
			mountains++
		} else if hike == 0 && step == 'D' {
			valleys++
		}
		if step == 'U' {
			hike++
		} else if step == 'D' {
			hike--
		}
	}

	return valleys
}
