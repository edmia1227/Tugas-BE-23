package main

import "fmt"

func encrypt(text string, shift int) string {

	var result string

	for _, char := range text {

		if char >= 'a' && char <= 'z' {

			result += string((char-'a'+rune(shift))%26 + 'a')

		} else if char >= 'A' && char <= 'Z' {

			result += string((char-'A'+rune(shift))%26 + 'A')

		} else {
			result += string(char)
		}
	}
	return result
}

func main() {

	shift := 1000
	input := "abcdefghijklmnopqrstuvwxyz"

	encryptText := encrypt(input, shift)

	fmt.Println("Input: ", input)
	fmt.Println("Output: ", encryptText)
}
