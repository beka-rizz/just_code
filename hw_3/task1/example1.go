package task1

import "fmt"

// 1. Reading from an empty channel
func Example1() {
	ch := make(chan int)
	fmt.Println(<-ch)
}
