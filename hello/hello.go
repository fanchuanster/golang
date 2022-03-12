package main

import (
	"fmt"
	// "hello/greetings"
	// "log"
	"math"
)

func Sqrt(x float64) float64 {
	// z := x/2
	z := 1
	// Newton's method - (z*z - x)/(2*z)
	// https://en.wikipedia.org/wiki/Newton%27s_method
	for ; math.Abs(z*z - x) > 0.00001; z -= (z*z - x)/(2*z) {
		fmt.Println(z)
	}
	return z
}

func FuzzBuzz(n int) {
	for i := 1; i <= 20; i++ {
		var res string
		if i % 3 == 0 {
			res += "fizz"
		}
		if i % 5 == 0 {
			if res != "" {
				res += " "
			}
			res += "buzz"
		}
		if res == "" {
			res = fmt.Sprintf("%v", i)
		}
		fmt.Println(res)	
	}
}

func main() {
	var n float64 = 200
	fmt.Println(Sqrt(n))
	fmt.Println("math.Sqrt is %f", math.Sqrt(n))
}
