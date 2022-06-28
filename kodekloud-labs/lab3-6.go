package main

import "fmt"

type Movie struct {
	name   string
	rating float32
}

func main() {
	mov := Movie{"xyz", 2.1}
	mov1 := Movie{"abc", 2.1}
	if mov.rating == mov1.rating || mov != mov1 {
		fmt.Println("condition met")
	} else if mov.rating == mov1.rating {
		fmt.Println("condition_2 met")
	}
}
