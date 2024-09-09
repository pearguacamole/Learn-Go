package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our Pizza Shop")
	fmt.Println("Enter a rating plz")

	reader := bufio.NewReader(os.Stdin)

	rating, _ := reader.ReadString('\n')
	numrating, err := strconv.ParseFloat(strings.TrimSpace(rating), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Add 1 to rating,", numrating+1)
	}

}
