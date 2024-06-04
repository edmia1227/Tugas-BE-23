package main

import (
	"fmt"
	"sort"
)

func dragonOfLowater(dragonHead, dragonKnight []int) int {

	sort.Ints(dragonHead)
	sort.Ints(dragonKnight)

	gold := 0
	KnIndex := 0

	for _, headSize := range dragonHead {

		for KnIndex < len(dragonKnight) && dragonKnight[KnIndex] < headSize {

			KnIndex++
		}

		if KnIndex == len(dragonKnight) {

			return -1
		}

		gold += dragonKnight[KnIndex]
		KnIndex++

	}
	return gold
}

func main() {

	heads := []int{7, 2}
	knights := []int{2, 1, 8, 5}
	gold := dragonOfLowater(heads, knights)

	if gold == -1 {

		fmt.Println("Knight Fall")

	} else {

		fmt.Println(gold)
	}
}
