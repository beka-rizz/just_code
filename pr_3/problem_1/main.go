package main

import "fmt"

func main() {
	nums := make(chan int)
	squares := make(chan int)

	go func() {
		//	sending on a channel "nums"
		defer close(nums)
		for _, num := range []int{1, 2, 3} {
			nums <- num
		}
	}()

	go func() {
		//	receiving from a channel and sending on a "squares" channel
		for num := range nums {
			squares <- num * num
		}
		close(squares)
	}()

	for num := range squares {
		fmt.Println(num)
	}
}
