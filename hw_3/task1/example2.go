package task1

import "fmt"

func writeData(c chan string) {
	c <- "Hello"
	close(c)
	c <- "World"
}

// 2. Sending on closed channel
func Example2() {
	ch := make(chan string)
	go writeData(ch)
	data := <-ch
	fmt.Println(data)
}
