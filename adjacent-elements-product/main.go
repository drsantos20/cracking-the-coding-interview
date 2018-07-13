package main

import (
	"fmt"
)

func main() {

	inputArray := []int{3, 6, -2, -5, 7, 3}

	max := -1001
	for i := 0; i < len(inputArray)-1; i++ {
		prod := inputArray[i] * inputArray[i+1]
		if prod > max {
			max = prod
		}
	}
	fmt.Println(inputArray)

}
