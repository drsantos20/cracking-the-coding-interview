package main

import (
	"fmt"
)

func main() {

	// Initialize a 2D array
	a := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	length := len(a) - 1

	for i := 0; i < len(a); i++ {

		for j := i; j < length-i; j++ {
			//first copy the top and swap the then
			temp := a[i][j]
			a[i][j] = a[length-j][i]
			a[length-j][i] = a[length-i][length-j]
			a[length-i][length-j] = a[j][length-i]
			a[j][length-i] = temp
		}
	}
	fmt.Println(a)

}
