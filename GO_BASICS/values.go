package main

import (
	"fmt"
	"math"
	"time"
)

const st = "constant"

func main() {
	fmt.Println("go" + "lang")
	fmt.Println("!+!=", 1+1)
	fmt.Println("9.0/4.0=", 9.0/4.0)
	fmt.Println("9.0/4.0=", 9.0/4)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
	var a string = "inital"
	fmt.Println(a)
	var b, c int = 13, 34
	fmt.Println(b, c)
	var d = true
	fmt.Println(d)
	var e int
	fmt.Println(e)
	var f bool
	fmt.Println(f)
	var g float32
	fmt.Println(g)
	h := "ggh"
	fmt.Println(h)
	var k string
	fmt.Println(k)
	start := time.Now()
	fmt.Println(st)
	const n = 500000000
	fmt.Println(n)
	const d1 = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d1))
	fmt.Println(math.Sin(n))
	fmt.Println(time.Since(start))

}
