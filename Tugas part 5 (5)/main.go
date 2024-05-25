package main

import "fmt"

func RemoveDuplikat(arr []int) int {

	if len(arr) == 0 {

		return 0
	}

	uniqindex := 1

	for i := 1; i < len(arr); i++ {

		if arr[i] != arr[i-1] {

			arr[uniqindex] = arr[i]
			uniqindex++
		}

	}

	return uniqindex
}

func main() {

	arr := []int{2, 3, 4, 5, 6, 9, 9}

	outputL := RemoveDuplikat(arr)

	fmt.Println("Input:", arr)
	fmt.Println("Output:", outputL)
}
