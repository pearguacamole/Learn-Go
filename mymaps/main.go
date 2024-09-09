package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to maps")

	languages := make(map[string]string)
	languages["rb"] = "ruby"
	languages["py"] = "python"
	languages["fe"] = "rust"
	languages["jv"] = "java"

	fmt.Println(languages)
	fmt.Println("rb shorts for", languages["rb"])

	fmt.Println("deleting rb")
	delete(languages, "rb")
	fmt.Println("map after deletion :", languages)

	// loops are intersting in go for maps
	for key, value := range languages {
		fmt.Printf("for key %v value is %v\n", key, value)
	}

}
