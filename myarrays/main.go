package main

import "fmt"

func main() {
	fmt.Println("welcome to array")
	var fruitlist [4]string
	fruitlist[0] = "apple"
	fruitlist[1] = "ape"
	//fruitlist[2] = "le"
	fruitlist[3] = "app"
	fmt.Println("list is ", fruitlist)
	fmt.Println("list size is ", len(fruitlist))
	var anotherList = [3]string{"a", "b", "c"}
	fmt.Println("list is ", anotherList)
	fmt.Println("list size is ", len(anotherList))

}
