package main

import "fmt"

func main() {
	fmt.Println("Welcome to my pointers")
	var ptr *int
	fmt.Println("default value of pointer is ", ptr)

	number := 23
	sideEye := &number
	fmt.Println("value inside pointer is ", sideEye)
	fmt.Println("valuse of obj its pointing to is ", *sideEye)

	fmt.Println("Playing wit it ")
	*sideEye = *sideEye * 2
	fmt.Println("the value inside number is", number)

}
