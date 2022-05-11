package main

import (
	"fmt"
	"math/rand"
	"time"
)

// waitForTask: In this pattern, the parent goroutine sends a signal to a
// child goroutine waiting to be told what to do.
func main() {
	ch := make(chan string)
	go func() {
		d := <-ch
		fmt.Println("child : recv'd signal :", d)
	}()
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	ch <- "data"
	fmt.Println("parent : sent signal")
	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}
