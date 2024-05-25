package main

import (
	"fmt"
	"strings"
)

func UbahHuruf(input string, mapping map[rune]rune) string {

	var result strings.Builder

	for _, char := range input {

		if newChar, ok := mapping[char]; ok {

			result.WriteRune(newChar)

		} else {
			result.WriteRune(char)
		}

	}

	return result.String()

}

func main() {

	input := "SEPULSA OKE"

	mapping := map[rune]rune{

		'A': 'k',
		'B': 'L',
		'C': 'M',
		'D': 'N',
		'E': 'O',
		'F': 'P',
		'G': 'Q',
		'H': 'R',
		'I': 'S',
		'J': 'T',
		'K': 'U',
		'L': 'V',
		'M': 'W',
		'N': 'X',
		'O': 'Y',
		'P': 'Z',
		'Q': 'A',
		'R': 'B',
		'S': 'C',
		'T': 'D',
		'U': 'E',
		'V': 'F',
		'W': 'G',
		'X': 'H',
		'Y': 'I',
		'Z': 'J',
	}

	Output := UbahHuruf(input, mapping)

	fmt.Println("input: ", input)
	fmt.Println("output: ", Output)

}
