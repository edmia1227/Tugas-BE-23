package main

import "fmt"

func binarySearch(arr []int, x int) int {

	left, right := 0, len(arr)-1

	for left <= right {

		mid := left + (right-left)/2

		if arr[mid] == x {
			return mid
		}
		if arr[mid] > x {

			right = mid - 1
		}

		if arr[mid] < x {

			left = mid + 1
		}
	}
	return -1
}

func main() {

	arr := []int{12, 15, 15, 19, 24, 31, 53, 59, 60}
	x := 53
	result := binarySearch(arr, x)

	if result != -1 {

		fmt.Println("Output: ", result)
	} else {
		fmt.Println("-1")
	}
}
