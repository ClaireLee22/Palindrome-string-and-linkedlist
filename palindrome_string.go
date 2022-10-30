package main

import "fmt"

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left += 1
		right -= 1
	}
	return true
}

func main() {
	s := "abccba"
	fmt.Println("Is palindrome?", isPalindrome(s))
}

// output: Is palindrome? true
