package main

import "fmt"

func bufferedCh1() {
	ch := make(chan string, 3)

	ch <- "ping"
	ch <- "pong"
	ch <- "dong"
	close(ch)

	// we need to close if we want to use such loop
	for str := range ch {
		fmt.Println(str)
	}
}

func bufferedCh2() {
	ch := make(chan string, 3)

	ch <- "ping"
	ch <- "pong"
	ch <- "dong"

	// we may not close if we know the exact amount of data be read from the channel
	for i := 0; i < 3; i++ {
		fmt.Println(<-ch)
	}
}
