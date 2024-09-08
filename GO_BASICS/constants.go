package main

import (
	"fmt"
	"math"
	"time"
)

const st = "constant"

func main() {
	start := time.Now()
	fmt.Println(st)
	const n = 500000000
	fmt.Println(n)
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
	fmt.Println(time.Since(start))

}
