package main

import (
	"fmt"
	"strings"
)

var finalAns float64

func fact(n float64, c chan float64) {
	f := 1.
	for i := 2.; i <= n; i++ {
		f *= i

	}
	c <- f
}
func anodaFact(n float64) {
	f := 1.
	for i := 2.; i <= n; i++ {
		f *= i

	}
	// return f
	finalAns = f
}

func main() {
	fmt.Println()
	fmt.Println(strings.Repeat("#-##-#", 10))
	fmt.Println("Without Channels")
	ch := make(chan float64)
	// defer close(ch)

	for i := 1.; i <= 50; i++ {
		// go fact(i,ch)
		go anodaFact(i)
		go fact(i, ch)

		f := <-ch
		// fmt.Println("Factorial of",i,"=",f)

		fmt.Println("Factorial of", i, "=", finalAns)
		fmt.Println("Using channels, the factorial of", i, "is:", f)
		//	Minor change to test out a signed commit
	}

}
