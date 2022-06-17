package main

import "fmt"

func main() {
	var fruits [3]string = [3]string{"apples", "banana", "lyche"}
	fmt.Println(fruits)

	var names [3]string = [3]string{"piyush", "peeyush", "peyush"}
	fmt.Println(names)
	fmt.Println(len(names))

	names[1] = "WarDaddy"
	fmt.Println(names[1])

	var grades [5]int = [5]int{10, 20, 30, 40, 50}
	for i := 0; i < len(grades); i++ {
		fmt.Println(grades[i])
	}

	for index, element := range grades {
		fmt.Println(index, "=>", element)
	}

}
