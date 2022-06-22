package main

import "fmt"

type Circle struct {
	x int
	y int
	r int
}

func main() {
	var c Circle
	c.x = 5
	c.y = 5
	c.r = 5
	fmt.Printf("%+v \n", c)
}
