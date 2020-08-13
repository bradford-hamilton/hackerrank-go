package main

// https://www.hackerrank.com/challenges/grading/problem
func gradingStudents(grades []int32) []int32 {
	graded := make([]int32, len(grades))
	for i, v := range grades {
		if v < 38 {
			graded[i] = v
		} else if v%5 > 2 {
			graded[i] = ((v / 5) + 1) * 5
		} else {
			graded[i] = v
		}
	}
	return graded
}
