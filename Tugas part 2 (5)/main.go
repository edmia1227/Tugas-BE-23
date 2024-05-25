package main

import (
	"fmt"
	"math"
)

func Eksponansi(base float64, eksponen float64) float64 {

	return math.Pow(base, eksponen)
}

func main() {

	var base, eksponen float64

	fmt.Print("ini: ")
	fmt.Scanln(&base)
	fmt.Print("dipangkat ini: ")
	fmt.Scanln(&eksponen)

	result := Eksponansi(base, eksponen)

	fmt.Print("Output ", result)
}
