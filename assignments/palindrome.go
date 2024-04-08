package main

import (
	"fmt"
	"strings"
)

func checkPalindrome(str string) bool {
	str = strings.ToLower(str)

	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	str := "malayalam"

	if checkPalindrome(str) {
		fmt.Println("Given string is a palindrome")
	} else {
		fmt.Println("Given string is not palindrome")
	}
}
