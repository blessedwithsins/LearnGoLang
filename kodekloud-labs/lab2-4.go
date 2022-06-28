package main

import "fmt"

type Employee struct {
	eid int
	id  int
}

func main() {
	employees := make([]Employee, 5)
	for i := range employees {
		employees[i] = Employee{i, i + 10}
		fmt.Println(employees[i])
	}
}
