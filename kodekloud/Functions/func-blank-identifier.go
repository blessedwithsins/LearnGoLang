package main

import "fmt"

func f() (int, int) {
	return 42, 53
}

func main() {
	a, b := f()
	fmt.Println(a, b)
	v, _ := f()
	fmt.Println(v)
}
