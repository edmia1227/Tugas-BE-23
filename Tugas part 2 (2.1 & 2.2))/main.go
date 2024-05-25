package main

import "fmt"

func FaktorBilangan(n int) []int {

	var faktor []int

	for i := 1; i <= n; i++ {

		if n%i == 0 {

			faktor = append(faktor, i)
		}

	}
	return faktor
}

func reverseFaktor(faktor []int) []int {

	reversed := make([]int, len(faktor))

	for i, v := range faktor {

		reversed[len(faktor)-1-i] = v
	}

	return reversed
}

func main() {

	var n int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scanln(&n)

	faktor := FaktorBilangan(n)
	fmt.Print("Output adalah ", n, faktor)

	reversedFaktor := reverseFaktor(faktor)
	fmt.Print("Output 2: ", n, reversedFaktor)
}
