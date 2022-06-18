package main

import "fmt"

func main() {
	codes := map[string]string{"en": "English", "fr": "French"}
	fmt.Println(codes)
	fmt.Println(len(codes))
	fmt.Println(codes["en"])

	code := make(map[string]string)
	fmt.Println(code)

	branch := map[string]string{"cs": "compsc", "it": "infotech"}
	value, found := branch["cs"]
	fmt.Println(found, value)
	branch["mech"] = "Mechanical"
	fmt.Println(branch)
	branch["cs"] = "Computer Science"
	delete(branch, "it")
	fmt.Println(branch)

	for key, value := range branch {
		fmt.Println(key, "===", value)
	}

}
