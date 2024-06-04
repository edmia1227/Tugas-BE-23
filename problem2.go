package main

import (
	"fmt"
	"sort"
)

func moneyChange(money int, denom []int) []int {

	sort.Sort(sort.Reverse(sort.IntSlice(denom)))

	change := []int{}

	for _, d := range denom {

		for money >= d {

			change = append(change, d)
			money -= d
		}
	}
	return change

}

func main() {

	money := 15321
	denom := []int{10000, 5000, 2000, 1000, 500, 200, 100, 50, 20, 10, 1}

	change := moneyChange(money, denom)

	fmt.Println("Kembalian: ", change)
}
