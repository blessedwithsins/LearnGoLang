package main

import "fmt"

func main() {
	arr := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	i := 2

	fmt.Println(arr)

	slice1 := arr[:i]
	slice2 := arr[i+1:]
	slice := append(slice1, slice2...)
	fmt.Println(slice)

}
