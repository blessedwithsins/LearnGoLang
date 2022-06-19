package main

import "fmt"

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	n := 5
	result := factorial(n)
	fmt.Println("Factorial of", n, "is :", result)
}
