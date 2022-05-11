package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

// fanOutSem: In this pattern, a semaphore is added to the fan out pattern
// to restrict the number of child goroutines that can be schedule to run.
func main() {
	// Number of goroutines
	children := 2000
	ch := make(chan string, children)

	g := runtime.GOMAXPROCS(0)

	// No of goroutines we would like in the running state
	sem := make(chan bool, g)

	// only g number of 2000 goroutines are in the running state, the rest are waiting in runnable
	for c := 0; c < children; c++ {
		go func(child int) {
			// signal in the sem ch till the buffer is full so only g no of goroutines can run at any given time
			sem <- true
			{
				t := time.Duration(rand.Intn(200)) * time.Millisecond
				time.Sleep(t)
				ch <- "data"
				fmt.Println("child : sent signal :", child)
			}
			// opens for another goroutine to do it's work
			<-sem
		}(c)
	}

	for children > 0 {
		d := <-ch
		children--
		fmt.Println(d)
		fmt.Println("parent : recv'd signal :", children)
	}

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}
