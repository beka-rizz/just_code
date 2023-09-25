package task1

func Example3() {
	messages := make(chan string)

	go func() {}()

	// Main goroutine is trying to send message to channel
	// But no other goroutines are running
	// And channel has no buffers
	// So it raises deadlock error
	messages <- "I wanna do something."
}
