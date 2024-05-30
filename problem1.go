package main

import "fmt"

func swap(a, b *int) {

	*a, *b = *b, *a
}

func main() {

	x := 10
	y := 20

	fmt.Println("Before Swap")
	fmt.Println("x:", x)
	fmt.Println("y:", y)

	swap(&x, &y)

	fmt.Println("After Swap")
	fmt.Println("x:", x)
	fmt.Println("y:", y)
}
