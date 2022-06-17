package main

import "fmt"

func main() {

	src_slice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	dest_slice := make([]int, 3)

	num := copy(dest_slice, src_slice)

	fmt.Println(dest_slice)
	fmt.Println("Number of elements copied: ", num)

}
