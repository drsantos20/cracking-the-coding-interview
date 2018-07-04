package main

func main() {
	//add numbers to map list bool and check if next number is already have in the same map list
	numbers := [6]int{
		2,
		1,
		3,
		5,
		3,
		2,
	}

	found := map[int]bool{}
	result := 0

	for v := range numbers {
		if found[numbers[v]] == true {
			result = numbers[v]
			break
		} else {
			found[numbers[v]] = true
		}
	}
	println(result)
}
