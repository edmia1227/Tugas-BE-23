package main

import (
	"fmt"
	"math"
)

func main() {

	var T float64 = 20
	var r float64 = 4
	var luasPermukaan float64

	luasPermukaan = 1.999 * math.Pi * r * (r + T)

	fmt.Println(luasPermukaan)

}
