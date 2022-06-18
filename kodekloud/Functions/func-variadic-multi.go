package main

import "fmt"

func printDetails(student string, subjects ...string) {
	fmt.Println("Hey", student, ", here are your subjects -")
	for _, sub := range subjects {
		fmt.Printf("%s, ", sub)
	}
}

func main() {
	printDetails("Joe", "Physics", "Chemistry")
}
