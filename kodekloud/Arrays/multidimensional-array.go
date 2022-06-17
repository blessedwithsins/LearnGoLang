package main

import "fmt"

func main() {
	arr := [3][2]int{{2, 4}, {4, 16}, {8, 64}}
	fmt.Println(arr)
	fmt.Println(arr[2][1])
}
