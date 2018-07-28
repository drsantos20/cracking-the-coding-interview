package main

func main() {

	statues := []int{
		1, 5, 2, 3,
		//6, 2, 3, 8,
		//5, 4, 6,
		//8, 1, 0, 4, 7,
		//4, 2, 7, 18,
	}

	//sort.Ints(statues)

	min := statues[0]
	max := statues[0]
	for i := 0; i < len(statues); i++ {
		if statues[i] < min {
			min = statues[i]
			println("new min ", min)
		}

		if statues[i] > max {
			max = statues[i]
			println("new max ", max)
		}

	}

	println("max ", max, " min ", min)

	a := max - min + 1 - len(statues)
	//return a

	print(a)

}
