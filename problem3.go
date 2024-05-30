package main

import "fmt"

type student struct {
	name  []string
	score []int
}

func (s student) Average() float64 {
	if len(s.score) == 0 {
		return 0
	}

	sum := 0
	for _, score := range s.score {
		sum += score
	}
	return float64(sum) / float64(len(s.score))

}

func (s student) min() int {

	if len(s.score) == 0 {

		return 0
	}

	min := s.score[0]
	for _, score := range s.score {

		if score < min {
			min = score
		}
	}
	return min
}

func (s student) max() int {

	if len(s.score) == 0 {

		return 0
	}

	max := s.score[0]
	for _, score := range s.score {

		if score > max {

			max = score
		}
	}
	return max
}

func main() {

	murid := student{
		name:  []string{"Rizky", "Kobar", "Ismail", "Umam", "Yopan"},
		score: []int{80, 75, 70, 75, 60},
	}

	fmt.Println("average score", murid.Average())
	fmt.Println("min score", murid.min())
	fmt.Println("max score", murid.max())
}
