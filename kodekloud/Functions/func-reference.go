package main

import "fmt"

func modify(s *string) {
	*s = "world"
}

func modify_slice(slice []int) {
	slice[0] = 100
}

func modify_maps(m map[string]int) {
	m["K"] = 75
}

func main() {
	a := "hello"
	fmt.Println(a)
	modify(&a)
	fmt.Println(a)

	slice := []int{10, 20, 30}
	fmt.Println(slice)
	modify_slice(slice)
	fmt.Println(slice)

	ascii_codes := make(map[string]int)
	ascii_codes["A"] = 65
	ascii_codes["F"] = 70
	fmt.Println(ascii_codes)
	modify_maps(ascii_codes)
	fmt.Println(ascii_codes)
}
