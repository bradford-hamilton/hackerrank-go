package main

func isPalindrome(str string) bool {
	for i := 0; i < len(str); i++ {
		if str[i] == str[len(str)-1-i] {
			continue
		} else {
			return false
		}
	}
	return true
}
