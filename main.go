package main

import (
	"fmt"
	works "homework_19/work"
	"sync"
)

func main() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	var waitG sync.WaitGroup

	for w := 1; w <= 3; w++ {
		waitG.Add(1)
		go works.Job(w, jobs, results, &waitG)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	go func() {
		waitG.Wait()
		close(results)
	}()

	for {
		select {
		case result, ok := <-results:
			if !ok {
				return
			}
			fmt.Println("Natija:", result)
		}
	}
}
