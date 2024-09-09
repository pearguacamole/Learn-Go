package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to slices")
	var lis = []string{"go", "python", "java"}
	fmt.Printf("type of lsit is %T\n", lis)

	lis = append(lis, "rust", "haskill", "assembly")
	fmt.Println(lis)

	lis = append(lis[1:3])
	fmt.Println(lis)

	highScore := make([]int, 4)
	highScore[0] = 34
	highScore[1] = 45
	highScore[2] = 78
	highScore[3] = 56
	highScore = append(highScore, 66, 23, 5)
	fmt.Println(highScore)
	sort.Ints(highScore)
	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore))

	//remove value form slices based on index
	var courses = []string{"react", "python", "java", "ruby", "rust"}
	fmt.Println(courses)
	var idx int = 2
	courses = append(courses[:idx], courses[idx+1:]...)
	fmt.Println(courses)

}
