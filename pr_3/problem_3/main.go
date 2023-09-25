package main

import "fmt"

var workks = make(map[int]int)

func main() {
	numJobs := 100
	numWorkers := 5

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for id := 1; id <= numWorkers; id++ {
		go worker(square, jobs, results, id)
	}

	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}
	close(jobs)

	for i := 1; i <= numJobs; i++ {
		fmt.Println(<-results)
	}
	close(results)

	for key, val := range workks {
		fmt.Println(key, "did", val, "jobs")
	}
}

func worker(f func(int) int, jobs <-chan int, results chan<- int, id int) {
	for num := range jobs {
		fmt.Printf("worker %d is in progress\n", id)
		workks[id]++
		results <- f(num)
	}
}

func square(num int) int {
	return num * num
}
