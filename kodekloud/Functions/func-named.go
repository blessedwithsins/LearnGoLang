package main

import "fmt"

func num(a int, b int) (sum int, diff int) {
	sum = a + b
	diff = a - b
	return
}

func main() {
	sum, difference := num(20, 10)
	fmt.Println(sum, " ", difference)
}
