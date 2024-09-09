package main

import (
	"fmt"
	"time"
)

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 is even")
	}
	if num := 9; num < 0 {
		fmt.Println("is neg")
	} else if num < 10 {
		fmt.Println("is 1 digit")
	} else {
		fmt.Println(num, "Hs many digits")
	}

	i1 := 2
	fmt.Print("Write ", i1, " as ")
	switch i1 {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}
	for i := range 3 {
		fmt.Println("range", i)
	}
	for {
		fmt.Println("disguise while")
		break
	}
	for n := range 10 {
		if n%2 == 0 {
			continue
		}
		fmt.Println("odd", n)
	}
}
