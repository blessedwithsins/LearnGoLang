package main

import "fmt"

func main() {
	x := func(l int, b int) int {
		return l * b
	}(20, 30)
	fmt.Printf("%T \n", x)
	fmt.Println(x)
}
