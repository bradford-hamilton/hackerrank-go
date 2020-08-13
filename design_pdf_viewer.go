package main

// https://www.hackerrank.com/challenges/designer-pdf-viewer/problem
// h: an array of integers representing the heights of each letter
// word: a string
func designerPdfViewer(h []int32, word string) int32 {
	// h - [1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 7]
	// word - zaba
	m := make(map[string]int32, len(h))
	j := 'a'
	for i := 0; i < len(h); i++ {
		m[string(j)] = h[i]
		j++
	}

	// Now in theory we have a map of letter -> height
	var largest int32
	for _, v := range word {
		if m[string(v)] > largest {
			largest = m[string(v)]
		}
	}
	return int32(len(word)) * largest
}
