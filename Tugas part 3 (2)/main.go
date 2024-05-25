package main

import "fmt"

func DrawXYZ(n int) {

	for i := 0; i < n; i++ {

		for j := 1; j <= n; j++ {

			ed := i*n + j

			if ed%3 == 0 {
				fmt.Print("X ")
			} else if ed%2 == 0 {
				fmt.Print("Z ")
			} else {
				fmt.Print("Y ")
			}

		}
		fmt.Println()

	}
}

func main() {

	n := 5

	DrawXYZ(n)

}
