package routine

// a goroutine is a lightweight thread managed by Go runtime.

import (
	"fmt"
	"time"
)

func say(s string) {
	fmt.Println(s, "======")
	for i:=0; i<3; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s, time.Millisecond)
	}
}

func TestGoroutine() {
	go say("Mrs")
	fmt.Println("hello goroutine")
}