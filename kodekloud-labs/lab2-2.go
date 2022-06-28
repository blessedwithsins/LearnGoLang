package main

import "fmt"

type Rectangle struct {
	length  int
	breadth int
}

func (r Rectangle) area() int {
	return r.length * r.breadth
}

func main() {
	r := Rectangle{breadth: 10, length: 5}
	fmt.Println(r.area())
	fmt.Println(r)
}
