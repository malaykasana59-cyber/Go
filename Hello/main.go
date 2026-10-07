package main

import (
	"fmt"
	"math"
	"math/rand"
)

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) { // func split(sum int) (int, int) {
	x = sum * 4 / 9 //     x := sum * 4 / 9
	y = sum - x     //	   y := sum - x
	return          //	   return x, y
} // }

var c, python, java bool

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println(math.Pi)
	fmt.Println(add(42, 13))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(17))

	var i int
	j := 0
	k := "10" // var k string = "10"
	fmt.Println(i, c, python, java, j, k)
}
