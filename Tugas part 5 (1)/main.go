package main

import (
	"fmt"
)

func compareStrings(a1, a2 string) string {

	if len(a1) < len(a2) {

		return a1
	}
	return a2

}

func main() {

	inputA := "AKA"
	inputB := "AKASHI"

	output := compareStrings(inputA, inputB)

	fmt.Println("Output:", output)
}
