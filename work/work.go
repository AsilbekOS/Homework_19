package works

import (
	"fmt"
	"sync"
	"time"
)

func Job(id int, jobs <-chan int, results chan<- int, waitG *sync.WaitGroup) {
	defer waitG.Done()
	for j := range jobs {
		fmt.Printf("Job %d is doing job %d\n", id, j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}
