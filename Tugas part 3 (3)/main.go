package main

import "fmt"

func TabelPerkalian(input int) {

	for i := 1; i <= input; i++ {
		for j := 1; j <= input; j++ {

			fmt.Printf("%4d", i*j)

		}
		fmt.Println()

	}
}

func main() {

	input := 9
	fmt.Printf("Output: ")

	TabelPerkalian(input)
}
