package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Beauty of Go Lang
// concurrency is misunderstood
// goroutines : Not scheduled by OS, Multiple Go routines on a single thread
// Wanna sound smart, Go uses Actor Model for Concurrency, Communicating Sequential processes(CSP)
// Go uses channels for Safe Communications
// Channels are buffered and unbuffered

func main() {
	// For running them Parallel, It's dangerous, Really really dangerous
	runtime.GOMAXPROCS(2)

	// waitGroup for the main function to wait
	var waitGrp sync.WaitGroup
	waitGrp.Add(2)

	// Adding go keyword before the func keyword turns it into a goroutine and the only one go routine would be sleeping
	// Self executing functions
	go func(){
		defer waitGrp.Done()
		// Entire program sleep
		time.Sleep(5 * time.Second)
		fmt.Println("Hello")
	}()

	go func(){
		defer waitGrp.Done()
		fmt.Println("Pluralsight")
	}()
	waitGrp.Wait()

}