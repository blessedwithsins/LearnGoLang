package main

import "fmt"

func addNum(a int, b int) int {
	sum := a + b
	return sum
}

func main() {
	addition := addNum(2, 3)
	fmt.Println(addition)
}
