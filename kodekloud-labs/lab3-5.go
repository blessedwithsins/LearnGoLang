package main

import "fmt"

type Movie struct {
	name   string
	rating float32
}

func getMovie(s string, r float32) (m Movie) {
	m = Movie{
		name:   s,
		rating: r,
	}
	return
}

func main() {
	mov := getMovie("xyz", 2.1)
	mov1 := getMovie("abc", 3.3)
	movies := make([]Movie, 5)
	movies = append(movies, mov)
	movies = append(movies, mov1)
	for _, value := range movies {
		fmt.Println(value)
	}
}
