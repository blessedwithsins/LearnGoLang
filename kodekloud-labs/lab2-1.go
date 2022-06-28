package main

import "fmt"

type Movie struct {
	name    string
	summary string
	rating  float32
}

func (m Movie) getSummary() {
	m.summary = "summary"
}

func (m *Movie) increaseRating() {
	m.rating += 1
}

func main() {
	mov := Movie{"xyz", "", 2.1}
	fmt.Printf("%+v", mov)
	mov.increaseRating()
	mov.getSummary()
	fmt.Printf("%+v", mov)
}
