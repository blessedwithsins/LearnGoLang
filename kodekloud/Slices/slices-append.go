package main

import "fmt"

func main() {
	arr := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	slice := arr[2:8]
	fmt.Println(slice)
	fmt.Println(cap(slice))
	fmt.Println(len(slice))

	slice = append(slice, 100, 300, 500)
	fmt.Println(slice)
	fmt.Println(cap(slice))
	fmt.Println(len(slice))

	array1 := [4]int{10, 20, 30, 40}
	slice1 := array1[1:3]
	array2 := [5]int{10, 20, 30, 40, 50}
	slice2 := array2[1:4]
	slice3 := append(slice1, slice2...)
	fmt.Println(slice3)

}
