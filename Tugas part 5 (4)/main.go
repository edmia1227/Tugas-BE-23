package main

import "fmt"

func maxSumSubArray(arr []int, k int) int {

	if len(arr) < k {
		return 0
	}

	sumMax := 0
	for i := 0; i < k; i++ {

		sumMax += arr[i]

	}

	sumCurrent := sumMax

	for i := k; i < len(arr); i++ {

		sumCurrent += arr[i] - arr[i-k]
		if sumCurrent > sumMax {

			sumMax = sumCurrent
		}

	}
	return sumMax
}

func main() {

	arr := []int{2, 3, 4, 1, 5}
	k := 2

	output := maxSumSubArray(arr, k)

	fmt.Println("Input:", arr)
	fmt.Println("k:", k)
	fmt.Println("Output:", output)
}
