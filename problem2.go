package main

import "fmt"

func getMinMax(numb []int, min *int, max *int) {

	if len(numb) == 0 {
		return
	}

	*min, *max = numb[0], numb[0]
	for _, value := range numb {

		if value < *min {
			*min = value
		}

		if value > *max {

			*max = value
		}
	}
}

func main() {

	numb := []int{1, 2, 3, 9, 7, 8}

	var min, max int

	getMinMax(numb, &min, &max)

	fmt.Println("Input:", numb)
	fmt.Println("minimum", min)
	fmt.Println("maximum", max)
}
