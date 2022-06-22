package main

import "fmt"

type s1 struct {
	x int
}

type s2 struct {
	x int
}

func main() {
	c := s1{x: 5}
	c1 := s1{x: 6}
	if c != c1 {
		fmt.Println("c & c1 have different values")
	}
	if c == c1 {
		fmt.Println("same values")
	}
}
