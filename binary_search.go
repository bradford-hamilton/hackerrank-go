package main

// must already be sorted
// returns index of found target, or -1
func binarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		midPoint := (low + high) / 2
		if arr[midPoint] < target {
			low = midPoint + 1
		} else {
			high = midPoint - 1
		}
	}

	if low == len(arr) || arr[low] != target {
		return -1
	}

	return low
}
