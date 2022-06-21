package main

import "fmt"

func main() {
	i := 10
	fmt.Printf("%T %v \n", &i, &i)
	fmt.Printf("%T %v \n", *(&i), *(&i))

	j := 10
	var ptr_j *int = &j
	k := "kilo"
	var ptr_k *string = &k
	fmt.Println(ptr_j)
	fmt.Println(ptr_k)

	s := "hello"
	var ptr_s *string = &s
	fmt.Println(ptr_s)
	var a = &s
	fmt.Println(a)
	b := &s
	fmt.Println(b)

	s1 := "hello"
	fmt.Printf("%T %v \n", s1, s1)
	ps1 := &s1
	*ps1 = "hello world"
	fmt.Printf("%T %v \n", s1, s1)

}
