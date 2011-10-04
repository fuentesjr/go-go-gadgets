package main

import "fmt"
import "time"
import "rand"
import "runtime"

func work(id int, comm_chan chan int) {
  secs := rand.Int63n(33) + 2
  fmt.Printf("worker %d working for %d seconds...\n", id, secs)
  time.Sleep(secs * 1000000000)
  fmt.Printf("worker %d done.\n", id)
  comm_chan <- 1
}

func main() {
  const NCPU = 2  // number of CPU cores
  runtime.GOMAXPROCS(NCPU)
  chan7 := make(chan int)

  // Start the workers
  fmt.Println("Spawning workers")
  for i := 0; i < NCPU; i++ {
    go work(i, chan7)
  }

  fmt.Println("Waiting for workers to complete")
  // Wait for workers to complete
  for i := 0; i < NCPU; i++ {
    <-chan7
  }

}
