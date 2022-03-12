package main

import (
	"fmt"
	// "hello/greetings"
	// "log"
	"math"
	"runtime"
	"time"
)

func ShowCurrentOSName() {
	fmt.Println("Go runs on")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Println(os)
	}
}

func SwitchWithNoCondition() {
	t := time.Now()
	defer fmt.Println("show at end of SwitchWithNoCondition")
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}
}

func Sqrt(x float64) float64 {
	// z := x/2
	z := float64(1)
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

func Test1() {
	t := time.Now()
	defer fmt.Println("return in the end, started at %v", t)
	var n float64 = 200
	fmt.Println(Sqrt(n))
	fmt.Println("math.Sqrt is %f", math.Sqrt(n))

	SwitchWithNoCondition()

	if 2>1 {
		// defer at function level, not scope level.
		// following defer is executed right before the first defer.
		// defer lines are executed in a reverse order
		// summary: defer lines are pushed to a stack, when function returns
		// executed at a first-in-last-out order - https://go.dev/tour/flowcontrol/13
		defer fmt.Println("defer ShowCurrentOSName")
		defer fmt.Println("defer ShowCurrentOSName2")
		ShowCurrentOSName()
	}
	
	fmt.Println("last line of main")
}

func Test2() {
	s := []int{1,2,3,4}
	printSlice(s)

	s = s[:2]
	printSlice(s)

	s = s[:3]
	printSlice(s)
}

func Test3() {
	s := make([]int, 0, 5)
	printSlice(s)

	s = append(s, 1,2,3,4,5)
	printSlice(s)

	s = append(s, 1,2,3)
	printSlice(s)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func TestFunctionClosures() {
	// each adder is stateful, each closure is bound to its own inner
	// variable sum.
	pos, neg := adder(), adder()
	for i:=1; i<5; i+=2 {
		fmt.Println(
			pos(i),
			neg(-i),
		)
	}

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	TestFunctionClosures()
}
