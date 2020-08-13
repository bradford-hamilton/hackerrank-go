package main

// keyboards: an array of integers representing keyboard prices
// drives: an array of integers representing drive prices
// b: the units of currency in Monica's budget
func getMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
	// var sortedKeyboards []int
	// for i := range keyboards {
	// 	sortedKeyboards[i] = int(keyboards[i])
	// }
	// var sortedDrives []int
	// for i := range drives {
	// 	sortedDrives[i] = int(drives[i])
	// }
	// sort.Ints(sortedKeyboards)
	// sort.Ints(sortedDrives)

	var max int32 = -1
	for i := 0; i < len(keyboards); i++ {
		j := i
		for ; j < len(drives); j++ {
			spend := keyboards[i] + drives[j]
			if (int32(spend) <= b) && int32(spend) > max {
				max = int32(spend)
			}
		}
	}

	return max
}
