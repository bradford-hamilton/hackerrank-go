package main

// get the largest height in heights
// return the difference between tallest height and k
func hurdleRace(k int32, height []int32) int32 {
	var tallest int32
	for _, v := range height {
		if v > tallest {
			tallest = v
		}
	}

	doses := tallest - k
	if doses < 1 {
		return 0
	}
	return doses
}
