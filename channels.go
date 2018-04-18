package main

import (
	"fmt"
	"time"
)

const (
	LOG_WORKER_ID bool = true
	WORK_LOAD     int  = 15
	NUM_WORKERS   int  = 5
)

type Work struct {
	x, y, z int
}

func log(msg string, id int) {
	if LOG_WORKER_ID {
		fmt.Println(msg, id)
	}
}

// Simple closure to generate worker id
func createWorkerId() func() int {
	id := 0
	return func() int {
		id++
		return id
	}
}

// Simple closure to track worker count
func manageWorkers() func() int {
	workerCount := NUM_WORKERS
	return func() int {
		workerCount--
		return workerCount
	}
}

// Use uni-directional channels
func worker(in <-chan *Work, out chan<- *Work, generateWorkerId func() int, manageWorkers func() int) {
	id := generateWorkerId()
	log("Creating worker with id: ", id)

	defer func() {
		if workersLeft := manageWorkers(); workersLeft == 0 {
			close(out)
		}
	}()

	for w := range in {
		log("Processing with worker id: ", id)
		w.z = w.x * w.y
		out <- w
	}
}

// Specify a send only channel
func doWork(in chan<- *Work) {
	defer close(in)
	for i := 0; i < WORK_LOAD; i++ {
		time.Sleep(200 * time.Millisecond)
		work := &Work{i, i, i}
		in <- work
	}
}

func receiveResult(r *Work) {
	fmt.Println("Result from worker: ", r.z)
}

func run() {
	in, out := make(chan *Work), make(chan *Work)
	workerIdGenerator := createWorkerId()
	workerManager := manageWorkers()
	for i := 0; i < NUM_WORKERS; i++ {
		go worker(in, out, workerIdGenerator, workerManager)
	}
	go doWork(in)
	for r := range out {
		receiveResult(r)
	}
}

func main() {
	run()
}
