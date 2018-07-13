package main

func main() {

	inputString := "aabaa"
	reverse := ""
	size := len(inputString)
	palindrome := false

	for j := 1; j <= len(inputString); j++ {
		reverse += string(inputString[size-j])
	}

	if reverse == inputString {
		palindrome = true
	}

	println(palindrome)

}
