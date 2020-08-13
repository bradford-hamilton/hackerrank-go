package main

func hurdleRace(k int32, height []int32) int32 {
	// get the largest height in heights
	var tallest int32
	for _, v := range height {
		if v > tallest {
			tallest = v
		}
	}

	// return the difference between tallest height and k
	doses := tallest - k
	if doses < 1 {
		return 0
	}

	return doses
}
