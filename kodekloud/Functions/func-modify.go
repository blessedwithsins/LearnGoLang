package main

import "fmt"

func modify(a string) {
	a = "world"
}

func main() {
	s := "hello"
	fmt.Println(s)
	modify(s)
	fmt.Println(s)
}
