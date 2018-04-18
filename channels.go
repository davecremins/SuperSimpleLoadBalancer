

package main

import (
   "fmt"
   "time"
)

type Work struct {
   x, y, z int
}

func worker(in <-chan *Work, out chan<- *Work) {
   for w := range in {
      w.z = w.x * w.y
      out <- w
   }
   close(out)
}

func sendLotsOfWork(in chan *Work){
   for i := 0; i < 10; i++ {
      time.Sleep(200 * time.Millisecond)
      var work = &Work{i, i, i}
      in <- work
   }
   close(in)
}

func receiveResult(out *Work){
   fmt.Println(out.z)
}

func Run() {
   in, out := make(chan *Work), make(chan *Work)
   go worker(in, out)
   go sendLotsOfWork(in)
   for {
      result, ok := <-out
      if !ok {
         break
      }
      receiveResult(result)
   }
}

func main() {
   Run()
}