package main

import "fmt"

type Employee struct {
	eid int
	id  int
}

func (e Employee) get_id() int {
	return e.eid + 10
}

func main() {
	employees := make([]Employee, 5)
	for i := range employees {
		employees[i] = Employee{eid: i}
		employees[i].id = employees[i].get_id()
		fmt.Printf("%+v\n", employees[i])
	}
}
