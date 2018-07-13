package main

import (
	"fmt"
	"math"
)

func main() {

	var highest_number float64

	a := []int{3, 6, -2, -5, 7, 3}
	max := len(a) - 1

	for i := 0; i < max; i++ {
		if i != max {
			highest_number = math.Abs(float64(a[i])) * math.Abs(float64(a[i+1]))

			if highest_number < math.Abs(float64(a[i]))*math.Abs(float64(a[i+1])) {
				highest_number = math.Abs(float64(a[i])) * math.Abs(float64(a[i+1]))
			}

		}

	}
	fmt.Println(highest_number)

}
