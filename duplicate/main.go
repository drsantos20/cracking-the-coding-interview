package main

import "fmt"

func removeDuplicates(elements []int) []int {
	// Use map to record duplicates as we find them.
	encountered := map[int]bool{}
	result := []int{}

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}

func main() {
	elements := []int{100, 200, 300, 100, 200, 400, 0}
	fmt.Println(elements)

	// Test our method.
	result := removeDuplicates(elements)
	fmt.Println(result)
}
