package main

import (
	"fmt"
	"math/rand"
	"time"
)

type work struct {
	id   int
	task string
}

type result struct {
	id   int
	task string
	res  string
}

func worker(w work, r chan<- result) {
	start := time.Now()
	fmt.Printf(" [%d] -> is working on task: '%s'\n", w.id, w.task)

	work := rand.Intn(5000)
	time.Sleep(time.Duration(work) * time.Millisecond)
	elapsed := time.Since(start)

	r <- result{
		id:   w.id,
		res:  fmt.Sprintf("computed in %s", elapsed),
		task: w.task,
	}
}

func dispatch(numWorkers int, res chan<- result) {
	for i := 0; i < numWorkers; i++ {
		fmt.Printf("[%d] task assigned to worker\n", i)
		go worker(work{
			id:   i,
			task: fmt.Sprintf("t__%d", rand.Int()),
		}, res)
	}
}

func main() {
	numWorkers := 10
	res := make(chan result, numWorkers)

	dispatch(numWorkers, res)

	for i := 0; i < numWorkers; i++ {
		r := <-res
		fmt.Printf("[%d] worker completed task '%s' with result '%s'!\n", r.id, r.task, r.res)
	}
	close(res)
}
