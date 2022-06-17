package main

import "fmt"

func main() {
	newslice := []int{10, 20, 30}
	fmt.Println(newslice)

	arr := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	slice1 := arr[2:8]
	fmt.Println(slice1)
	fmt.Println(cap(arr))
	fmt.Println(len(arr))

	slice1[2] = 200

	fmt.Println("After Modification")
	fmt.Println(arr)
	fmt.Println(slice1)

	subslice := slice1[1:3]
	fmt.Println(subslice)

	slice2 := make([]int, 5, 8)
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

}
