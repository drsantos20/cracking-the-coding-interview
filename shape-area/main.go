package main

func main() {

	shapeArea(5)

}

func shapeArea(n int) int {

	if n > 1 {
		return 4*(n-1) + shapeArea(n-1)
	}

	return 1

	//aqui é feito uma chamada recursiva, onde o há uma area comum de 4 lados, a chamada é feita até n ser igual a 1

	//1 5 13 25
	//3: shapeArea we had at 2 + (2 * 4)

}
