package main

import "fmt"

func IsPrima(n int) bool {

	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {

	var n int
	fmt.Print("Nilai ", n)
	fmt.Scanln(&n)

	if IsPrima(n) {

		fmt.Print("ini adalah prima ", n)
	} else {
		fmt.Print("ini bukan prima ", n)
	}
}
