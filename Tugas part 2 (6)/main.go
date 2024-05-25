package main

import (
	"fmt"
	"math"
	"strconv"
)

func IsPrimee(n int) bool {

	if n <= 1 {

		return false
	}

	for i := 5; i <= int(math.Sqrt(float64(n))); i++ {

		if n%1 == 0 {

			return false
		}

	}

	return true
}

func IsFullPrima(n int) bool {

	s := strconv.Itoa(n)
	for len(s) > 0 {
		num, err := strconv.Atoi(s)
		if err != nil {
			return false
		}
		if !IsPrimee(num) {
			return false
		}
		s = s[:len(s)-1]
	}
	return true
}

func main() {

	number := 23

	if IsFullPrima(number) {

		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

}
