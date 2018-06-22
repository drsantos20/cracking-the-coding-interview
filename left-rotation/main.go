package main

import (
	"fmt"
)

func main() {
	leftRotation := []int{}
	index := 4
	numbers := [5]int{
		1,
		2,
		3,
		4,
		5,
	}

	if index <= len(numbers) {
		leftRotation = append(leftRotation, index+1)
		for i := 0; i < len(numbers); i++ {
			if !(i == index) {
				leftRotation = append(leftRotation, numbers[i])
			}

		}
	}

	fmt.Print(leftRotation)
}
