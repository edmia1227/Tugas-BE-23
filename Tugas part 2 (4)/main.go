package main

import (
	"fmt"
	"strings"
)

func IsPalindrome(s string) bool {

	s = strings.ToLower(strings.ReplaceAll(s, " ", ""))

	length := len(s)

	for i := 0; i < length/2; i++ {
		if s[i] != s[length-1-i] {
			return false
		}

	}

	return true
}

func main() {

	var input string
	fmt.Print("masukkan kata: ")
	fmt.Scanln(&input)

	if IsPalindrome(input) {

		fmt.Println("ini adalah palindrome")
	} else {
		fmt.Println("ini bukan palindrome")
	}
}
