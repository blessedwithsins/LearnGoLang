package main

import "fmt"

func main() {

	arr := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	for index, value := range arr {
		fmt.Println(index, "=>", value)
	}

	for _, value := range arr {
		fmt.Println(value)
	}

}
