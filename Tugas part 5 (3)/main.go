package main

import (
	"fmt"
)

func ArrayUniq(arr1, arr2 []int) []int {

	cek := make(map[int]bool)

	for _, value := range arr2 {

		cek[value] = true
	}

	var result []int
	for _, value := range arr1 {

		if !cek[value] {
			result = append(result, value)
		}
	}

	return result

}

func main() {

	arr1 := []int{1, 2, 3, 4}
	arr2 := []int{1, 3, 5, 10, 16}

	output := ArrayUniq(arr1, arr2)

	fmt.Println("input: ")
	fmt.Println("arr1:", arr1)
	fmt.Println("arr2:", arr2)
	fmt.Println("Output:", output)

}
