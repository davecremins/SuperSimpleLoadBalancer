package main

import (
	"fmt"
	"time"
)

const LOG_WORKER_ID bool = false

type Work struct {
	x, y, z int
}

func log(msg string, id int) {
	if LOG_WORKER_ID {
		fmt.Println(msg, id)
	}
}

// Use uni-directional channels
func worker(in <-chan *Work, out chan<- *Work, id int) {
	log("creating worker with id: ", id)
	for w := range in {
		log("processing with worker id: ", id)
		w.z = w.x * w.y
		out <- w
	}
	close(out)
}

// Specify a send only channel
func sendLotsOfWork(in chan<- *Work) {
	for i := 0; i < 10; i++ {
		time.Sleep(200 * time.Millisecond)
		var work = &Work{i, i, i}
		in <- work
	}
	// time.Sleep(2000 * time.Millisecond)
	close(in)
}

func receiveResult(r *Work) {
	fmt.Println("Result from worker: ", r.z)
}

func Run() {
	in, out := make(chan *Work), make(chan *Work)
	NumWorkers := 3
	for i := 0; i < NumWorkers; i++ {
		go worker(in, out, i)
	}
	go sendLotsOfWork(in)
	for r := range out {
		receiveResult(r)
	}
}

func main() {
	Run()
}
