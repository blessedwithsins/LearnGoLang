package main

import "fmt"

type Student struct {
	name   string
	rollNo int
	marks  []int
	grades map[string]int
}

func main() {
	st := Student{
		name:   "Piyush",
		rollNo: 10,
	}
	fmt.Printf("%+v \n", st)
}
