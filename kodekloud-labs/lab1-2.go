package main

import "fmt"

type Student interface {
	getPercentage() int
	getName() string
}

type Undergrad struct {
	name   string
	grades []int
}

func (u Undergrad) getPercentage() int {
	sum := 0
	for _, v := range u.grades {
		sum += v
	}
	return sum / len(u.grades)
}
func (u Undergrad) getName() string {
	return u.name
}

func printData(s Student) {
	fmt.Println(s.getName())
	fmt.Println(s.getPercentage())
}

func main() {
	grades := []int{90, 75, 80}
	u := Undergrad{"Ross", grades}
	printData(u)
}
