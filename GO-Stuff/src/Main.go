package main

import (
	"fmt"
	"math"
)


func main() {
	fmt.Println("Hello, playground")
	x := 5.32
	result := math.Pow(x, 8)
	fmt.Println(x, "to the power of ", 8, " = ", result)
	var balance = 10.0
	_ = balance
}
