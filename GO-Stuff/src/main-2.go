package main

import "fmt"

func SayHello() {
	fmt.Println("Hello package v1.0.0!")
	fmt.Println("Nothing is brokey")

	var x int = 10
	_ = x
}
func main() {
	SayHello()
}
