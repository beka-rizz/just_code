package main

import (
	"fmt"
	"sync"
)

func main() {
	a := make(chan int)
	b := make(chan int)
	c := make(chan int)

	go func() {
		for _, num := range []int{1, 2, 3, 4} {
			a <- num
		}
		close(a)
	}()

	go func() {
		for _, num := range []int{5, 6} {
			b <- num
		}
		close(b)
	}()

	go func() {
		for _, num := range []int{7, 8, 9} {
			c <- num
		}
		close(c)
	}()

	for num := range mergeChannels(a, b, c) {
		fmt.Println(num)
	}
}

func mergeChannels(channels ...chan int) chan int {
	result := make(chan int)

	go func() {
		wg := &sync.WaitGroup{}
		wg.Add(len(channels))
		for _, channel := range channels {
			go func(ch_ chan int) {
				for num := range ch_ {
					result <- num
				}
				wg.Done()
			}(channel)
		}
		wg.Wait()
		close(result)
	}()

	return result
}
